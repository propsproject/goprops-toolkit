package logging

import (
	"os"

	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"
	"fmt"
	"strings"
	"runtime"
)

const (
	devLogOut = "./logs/development.log"
	stderr    = "stderr"
)

type PropsLogger interface {
	Info(v ...interface{})
	Warn(v ...interface{})
	Error(v ...interface{})
	Fatal(v ...interface{})
}

type Message interface {
	SetMessage(v ...interface{}) *Message
	InitErr(err error, logger *PropsLogger) *Message
	WithField(field ...Field) *Message
	Log()
}

type Log struct {
	loggerInstance *PLogger
	severity int // 0=info, 1=warn, 2=error, 3=fatal
	message string
	Fields []Field
}

func (l *Log) WithField(k string, v interface{}) *Log {
	l.Fields = append(l.Fields, NewField(k, v))
	return l
}

func (l *Log) setMessage(v interface{}) *Log {
	l.message = GetString(v)
	return l
}

func (l *Log) Log() {
	l.loggerInstance.log(l.message, l.severity, l.Fields...)
}

func NewLog(logger *PLogger, severity int, v interface{}) *Log {
	log := &Log{loggerInstance: logger, severity: severity}
	return log.setMessage(v)
}

// PLogger light convience wrapper around zap logging
type PLogger struct {
	zapLogger *zap.Logger
	DefaultFields []Field
}

// Field convience struct for logs with fields, (support for strings only so far)
type Field struct {
	Key   string
	Value string
}

// Logger ...
var Logger *PLogger

// Info ...
func (l *PLogger) Info(v ...interface{}) *Log {
	return NewLog(&PLogger{zapLogger: l.zapLogger.WithOptions()}, 0, v)
}

//Warn ...
func (l *PLogger) Warn(v ...interface{}) *Log {
	return NewLog(&PLogger{zapLogger: l.zapLogger.WithOptions()}, 1, v)
}

// Error ...
func (l *PLogger) Error(v ...interface{}) *Log {
	return NewLog(&PLogger{zapLogger: l.zapLogger.WithOptions()}, 2, v)
}

// Fatal ...
func (l *PLogger) Fatal(v ...interface{}) *Log {
	return NewLog(&PLogger{zapLogger: l.zapLogger.WithOptions()}, 3, v)
}

func (l *PLogger) log(msg string, severity int, fields ...Field) {
	allFields := append(l.DefaultFields, fields...)
	switch severity {
	case 0:
		l.zapLogger.Info(msg, ToZapFields(allFields...)...)
	case 1:
		l.zapLogger.Warn(msg, ToZapFields(allFields...)...)
	case 2:
		l.zapLogger.Error(msg, ToZapFields(allFields...)...)
	case 3:
		l.zapLogger.Fatal(msg, ToZapFields(allFields...)...)
		os.Exit(1)
	}
}

func (l *PLogger) Instance(name string) *PLogger {
	return &PLogger{
		zapLogger: l.zapLogger.Named(name),
		DefaultFields: l.DefaultFields,
	}
}

// ToZapFields creates fields map for log message
func ToZapFields(f ...Field) []zapcore.Field {
	var fields []zapcore.Field
	for _, field := range f {
		fields = append(fields, zap.String(field.Key, field.Value))
	}

	return fields
}

func NewField(k string, v interface{}) Field {
	return Field{Key:k, Value: GetString(v)}
}

func GetString(v interface{}) string {
	var builder strings.Builder
	for _, value := range v.([]interface{}) {
		switch value.(type) {
		case string:
			builder.WriteString(fmt.Sprintf("%s ", value.(string)))
		case error:
			builder.WriteString(fmt.Sprintf("%s ", value.(error).Error()))
		default:
			builder.WriteString(fmt.Sprintf("%v ", value))
		}
	}

	return builder.String()
}

func GetCallerExtra() []Field {
	pc := make([]uintptr, 15)
	n := runtime.Callers(15, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()

	return []Field{
		NewField("file-name", frame.File),
		NewField("line-number", frame.Line),
		NewField("func-name", frame.Function),
	}
}

var highPriority = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
	return lvl >= zapcore.DebugLevel
})
var lowPriority = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
	return lvl < zapcore.ErrorLevel
})

var (
	jsonDebugging = zapcore.AddSync(os.Stdout)
	jsonErrors = zapcore.AddSync(os.Stdout)
	productionConfig = zap.NewProductionEncoderConfig()

	developmentConfig = zap.NewDevelopmentConfig()
)

func WithPID() []zap.Option {
	return []zap.Option{zap.Fields(zap.Int("pid", os.Getpid()))}
}

func WithServiceName(name string) []zap.Option {
	return []zap.Option{zap.Fields(zap.String("service", name))}
}

func WithLineNumber() []zap.Option  {
	return []zap.Option{zap.AddCaller(), zap.AddCallerSkip(2)}
}

func DefaultOptions(name string) []zap.Option {
	opts :=  append(WithPID(), WithLineNumber()...,)
	return append(opts, WithServiceName(name)...)
}

func initProductionLogger(service string) *zap.Logger {
	jsonEncoder := zapcore.NewJSONEncoder(productionConfig)
	cors := []zapcore.Core{zapcore.NewCore(jsonEncoder, jsonErrors, highPriority), zapcore.NewCore(jsonEncoder, jsonDebugging, highPriority)}
	core := zapcore.NewTee(cors...)
	return zap.New(core).WithOptions(DefaultOptions(service)...)
}

func initDevelopmentLogger(service string) *zap.Logger {
	developmentConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	l, _:= developmentConfig.Build(DefaultOptions(service)...)
	return l
}

func NewLogger(service string, production bool, defaultFields ...Field) *PLogger {
	var zapLogger *zap.Logger

	if production {
		zapLogger = initProductionLogger(service)
	} else {
		zapLogger = initDevelopmentLogger(service)
	}

	return &PLogger{
		zapLogger: zapLogger,
		DefaultFields: defaultFields,
	}
}
