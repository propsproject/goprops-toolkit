package v2

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

// LoggerWrapper light convience wrapper around zap logger
type LoggerWrapper struct {
	zapLogger *zap.Logger
}

// Field convience struct for logs with fields, (support for strings only so far)
type Field struct {
	Key   string
	Value string
}

// Logger ...
var Logger *LoggerWrapper

// Info ...
func (l *LoggerWrapper) Info(msg string, data ...Field) {
	l.zapLogger.Info(msg, Fields(data...)...)
}

// Warn ...
func (l *LoggerWrapper) Warn(msg string, data ...Field) {
	l.zapLogger.Warn(msg, Fields(data...)...)
}

// Error ...
func (l *LoggerWrapper) Error(err error, data ...Field) {
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

func init() {
	jsonDebugging := zapcore.AddSync(ioutil.Discard)
	jsonErrors := zapcore.AddSync(ioutil.Discard)
	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)

	jsonEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	core := zapcore.NewTee(
		zapcore.NewCore(jsonEncoder, jsonErrors, highPriority),
		zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
		zapcore.NewCore(jsonEncoder, jsonDebugging, lowPriority),
		zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority),
	)

	// var opts []zap.Option
	// opts = append(opts,
	// 	zap.Fields(zap.Int("pid", os.Getpid())),
	// 	zap.Fields(zap.String("exe", path.Base(os.Args[0]))),
	// )

	Logger = &LoggerWrapper{
		zapLogger: zap.New(core),
	}
	defer Logger.zapLogger.Sync()
}
