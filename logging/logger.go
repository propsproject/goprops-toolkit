package logging

import (
	"os"

	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"
	"fmt"
	"strings"
	"runtime"
)

type Logger struct {
	zapLogger *zap.Logger
	Fields []Field
	DefaultFields []Field
}

func (l *Logger) WithField(k string, v ...interface{}) *Logger {
	l.Fields = append(l.Fields, NewField(k, v))
	return l
}

func (l *Logger) flushFields() {
	l.Fields = []Field{}
}

// Field convience struct for logs with fields, (support for strings only so far)
type Field struct {
	Key   string
	Value string
}

// Info ...
func (l *Logger) Info(v ...interface{}) {
	l.zapLogger.Info(fmt.Sprintf("%v", v...), ToZapFields(l.DefaultFields...)...)
	l.flushFields()
}

//Warn ...
func (l *Logger) Warn(v ...interface{}) {
	l.zapLogger.Warn(fmt.Sprintf("%v", v...), ToZapFields(l.DefaultFields...)...)
	l.flushFields()
}

//Error ...
func (l *Logger) Error(v ...interface{}) {
	l.zapLogger.Error(fmt.Sprintf("%v", v...), ToZapFields(l.DefaultFields...)...)
	l.flushFields()
}

// Fatal ...
func (l *Logger) Fatal(v ...interface{}) {
	l.zapLogger.Error(fmt.Sprintf("%v", v...), ToZapFields(l.DefaultFields...)...)
	os.Exit(1)
	l.flushFields()
}


// Infof ...
func (l *Logger) Infof(str string, v ...interface{}) {
	l.zapLogger.Info(fmt.Sprintf(str, v...), ToZapFields(l.DefaultFields...)...)
	l.flushFields()
}

//Warnf ...
func (l *Logger) Warnf(str string, v ...interface{}) {
	l.zapLogger.Warn(fmt.Sprintf(str, v...), ToZapFields(l.DefaultFields...)...)
	l.flushFields()
}

//Errorf ...
func (l *Logger) Errorf(str string, v ...interface{}) {
	l.zapLogger.Error(fmt.Sprintf(str, v...), ToZapFields(l.DefaultFields...)...)
	l.flushFields()
}

// Fatalf ...
func (l *Logger) Fatalf(str string, v ...interface{}) {
	l.zapLogger.Error(fmt.Sprintf(str, v...), ToZapFields(l.DefaultFields...)...)
	os.Exit(1)
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
	return Field{Key:k, Value: fmt.Sprintf("%v", v)}
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

func Get(service string, production bool, defaultFields ...Field) *Logger {
	var zapLogger *zap.Logger

	if production {
		zapLogger = initProductionLogger(service)
	} else {
		zapLogger = initDevelopmentLogger(service)
	}

	return &Logger{
		zapLogger: zapLogger,
		DefaultFields: defaultFields,
	}
}
