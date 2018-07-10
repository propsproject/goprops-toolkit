package logger

import "go.uber.org/zap/zapcore"

type Plugin interface {
	WriteSyncer() zapcore.WriteSyncer
	Flush()
	Write(p []byte) (n int, err error)
}