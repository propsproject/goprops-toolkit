package v2

import (
	"fmt"
	"os"
	"path"

	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"
)

const (
	devLogOut = "./logs/development.log"
	stderr    = "stderr"
)

// LoggerWrapper light convience wrapper around zap logger
type LoggerWrapper struct {
	zapLogger *zap.SugaredLogger
}

// Field convience struct for logs with fields
type Field struct {
	Key   string
	Value interface{}
}

// Logger ...
var Logger *LoggerWrapper

// Info ...
func (l *LoggerWrapper) Info(msg string, data map[string]interface{}) {
	if args, ok := l.unpackArgs(data); ok {
		l.zapLogger.Infof(msg, args)
		return
	}
	l.zapLogger.Info(msg)
}

// Warn ...
func (l *LoggerWrapper) Warn(msg string, data map[string]interface{}) {
	if args, ok := l.unpackArgs(data); ok {
		l.zapLogger.Warnf(msg, args)
		return
	}
	l.zapLogger.Warn(msg)
}

// Error ...
func (l *LoggerWrapper) Error(err error, data map[string]interface{}) {
	if args, ok := l.unpackArgs(data); ok {
		l.zapLogger.Errorf(err.Error(), args)
		return
	}
	l.zapLogger.Error(err.Error())
}

func (l *LoggerWrapper) unpackArgs(args map[string]interface{}) ([]interface{}, bool) {
	var unpacked []interface{}
	for key, val := range args {
		unpacked = append(unpacked, key, val)
	}
	return unpacked, len(unpacked) > 0
}

// Fields creates fields map for log message
func Fields(f ...Field) map[string]interface{} {
	fields := make(map[string]interface{})
	for _, val := range f {
		fields[val.Key] = val.Value
	}

	return fields
}

func init() {
	cfg := zap.NewDevelopmentConfig()
	cfg.OutputPaths = []string{devLogOut, stderr}
	cfg.DisableCaller = true
	cfg.DisableStacktrace = true
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	initialFields := make(map[string]interface{})
	initialFields["PID"] = os.Getpid()
	initialFields["EXE"] = path.Base(os.Args[0])
	cfg.InitialFields = initialFields

	logger, err := cfg.Build()
	if err != nil {
		panic(fmt.Errorf("Could not initialize fallback logger: %v", err))
	}

	Logger = &LoggerWrapper{logger.Sugar()}
}
