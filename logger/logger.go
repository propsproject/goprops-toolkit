package logger

import (
	"sync"

	"go.uber.org/zap"
)

var instance *zap.SugaredLogger
var once sync.Once

func sharedLogger() *zap.SugaredLogger {
	once.Do(func() {
		cfg := zap.NewProductionConfig()
		cfg.OutputPaths = []string{
			"./logs/development.log",
			"stderr",
		}
		cfg.DisableCaller = true
		cfg.DisableStacktrace = true

		logger, _ := cfg.Build()
		instance = logger.Sugar()
	})
	return instance
}

func Init() {
	Info("Congiguring Logger")
}

func Info(msg string) {
	sharedLogger().Infow(msg)
}

func Error(err error) {
	sharedLogger().Errorw(err.Error())
}
