package logging

import (
	"os"

	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"
	"strings"
	"fmt"
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

func (l *Log) WithField(k, v string) *Log {
	l.Fields = append(l.Fields, Field{Key:k, Value:v})
	return l
}

func (l *Log) setMessage(v ...interface{}) *Log {
	var builder strings.Builder
	for _, str := range v {
		builder.WriteString(fmt.Sprintf("%v", str))
	}
	l.message = builder.String()
	return l
}

func (l *Log) Log() {
	l.loggerInstance.log(l.message, l.severity, l.Fields...)
}

func NewLog(logger *PLogger, severity int, v ...interface{}) *Log {
	log := &Log{loggerInstance: logger, severity: severity}
	return log.setMessage(v)
}

// PLogger light convience wrapper around zap logging
type PLogger struct {
	zapLogger *zap.Logger
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
	return NewLog(l, 0, v...)
}

//Warn ...
func (l *PLogger) Warn(v ...interface{}) *Log {
	return NewLog(l, 1, v...)
}

// Error ...
func (l *PLogger) Error(v ...interface{}) *Log {
	return NewLog(l, 2, v...)
}

// Fatal ...
func (l *PLogger) Fatal(v ...interface{}) *Log {
	return NewLog(l, 3, v...)
}

func (l *PLogger) log(msg string, severity int, fields ...Field) {
	switch severity {
	case 0:
		l.zapLogger.Info(msg, Fields(fields...)...)
	case 1:
		l.zapLogger.Warn(msg, Fields(fields...)...)
	case 2:
		l.zapLogger.Error(msg, Fields(fields...)...)
	case 3:
		l.zapLogger.Fatal(msg, Fields(fields...)...)
		os.Exit(1)
	}
}

// Fields creates fields map for log message
func Fields(f ...Field) []zapcore.Field {
	var fields []zapcore.Field
	for _, field := range f {
		fields = append(fields, zap.String(field.Key, field.Value))
	}

	return fields
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

	consoleDebugging = zapcore.Lock(os.Stdout)
	consoleErrors = zapcore.Lock(os.Stderr)
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
	productionConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	jsonEncoder := zapcore.NewJSONEncoder(productionConfig)
	cors := []zapcore.Core{zapcore.NewCore(jsonEncoder, jsonErrors, highPriority), zapcore.NewCore(jsonEncoder, jsonDebugging, lowPriority)}
	core := zapcore.NewTee(cors...)
	return zap.New(core).WithOptions(DefaultOptions(service)...)
}

func initDevelopmentLogger(service string) *zap.Logger {
	developmentConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(developmentConfig.EncoderConfig)
	cors := []zapcore.Core{zapcore.NewCore(consoleEncoder, consoleErrors, highPriority), zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority)}
	core := zapcore.NewTee(cors...)
	return zap.New(core).WithOptions(DefaultOptions(service)...)
}

func NewLogger(service string, production bool) *PLogger {
	var zapLogger *zap.Logger
	if production {
		zapLogger = initProductionLogger(service)
	} else {
		zapLogger = initDevelopmentLogger(service)
	}

	return &PLogger{zapLogger}
}
