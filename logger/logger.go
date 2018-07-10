package logger

import (
	"io/ioutil"
	"os"

	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"
)

const (
	devLogOut = "./logs/development.log"
	stderr    = "stderr"
)

// Wrapper light convience wrapper around zap logger
type Wrapper struct {
	zapLogger *zap.Logger
}

// Field convience struct for logs with fields, (support for strings only so far)
type Field struct {
	Key   string
	Value string
}

// Logger ...
var Logger *Wrapper

// Info ...
func (l *Wrapper) Info(msg string, data ...Field) {
	l.zapLogger.Info(msg, Fields(data...)...)
}

// Warn ...
func (l *Wrapper) Warn(msg string, data ...Field) {
	l.zapLogger.Warn(msg, Fields(data...)...)
}

// Error ...
func (l *Wrapper) Error(err error, data ...Field) {
	l.zapLogger.Error(err.Error(), Fields(data...)...)
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
	return lvl >= zapcore.ErrorLevel
})
var lowPriority = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
	return lvl < zapcore.ErrorLevel
})

func NewLogger(plugins ...Plugin) *Wrapper {
	jsonDebugging := zapcore.AddSync(ioutil.Discard)
	jsonErrors := zapcore.AddSync(ioutil.Discard)
	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)

	jsonEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(config.EncoderConfig)

	cors := []zapcore.Core{zapcore.NewCore(jsonEncoder, jsonErrors, highPriority),
		zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
		zapcore.NewCore(jsonEncoder, jsonDebugging, lowPriority),
		zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority)}

	if len(plugins) > 0 {
		for _, p := range plugins {
			cor := zapcore.NewCore(jsonEncoder, zapcore.Lock(p.WriteSyncer()), highPriority)
			cors = append(cors, cor)
		}
	}

	core := zapcore.NewTee(cors...)

	// var opts []zap.Option
	// opts = append(opts,
	// 	zap.Fields(zap.Int("pid", os.Getpid())),
	// 	zap.Fields(zap.String("exe", path.Base(os.Args[0]))),
	// )
	return &Wrapper{
		zapLogger: zap.New(core),
	}
}
