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
func (l *LoggerWrapper) Error(msg string, data map[string]interface{}) {
	if args, ok := l.unpackArgs(data); ok {
		l.zapLogger.Errorf(msg, args)
		return
	}
	l.zapLogger.Error(msg)
}

func (l *LoggerWrapper) unpackArgs(args map[string]interface{}) ([]interface{}, bool) {
	var unpacked []interface{}
	for key, val := range args {
		unpacked = append(unpacked, key, val)
	}
	return unpacked, len(unpacked) > 0
}

func init() {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{devLogOut, stderr}
	cfg.DisableCaller = true
	cfg.DisableStacktrace = true
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	var initialFields map[string]interface{}
	initialFields["PID"] = os.Getpid()
	initialFields["EXE"] = path.Base(os.Args[0])
	cfg.InitialFields = initialFields

	logger, err := cfg.Build()
	if err != nil {
		panic(fmt.Errorf("Could not initialize fallback logger: %v", err))
	}

	Logger = &LoggerWrapper{logger.Sugar()}
}
