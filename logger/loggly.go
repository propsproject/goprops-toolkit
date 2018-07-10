package logger

import (
	"github.com/segmentio/go-loggly"
	"io"
	"go.uber.org/zap/zapcore"
)

//LogglyWrapper . . .
type LogglyWrapper struct {
	client *loggly.Client
	io.Writer
}

func New(token string, tags ...string) *LogglyWrapper {
	client := loggly.New(token, tags...)
	client.Defaults = loggly.Message{}

	return &LogglyWrapper{client: client}
}

func (l *LogglyWrapper) WriteSyncer() zapcore.WriteSyncer {
	return zapcore.AddSync(l)
}

func (l *LogglyWrapper) Flush() {
	l.client.Flush()
}

func (z *LogglyWrapper) Write(p []byte) (n int, err error) {
	return z.client.Write(p)
}