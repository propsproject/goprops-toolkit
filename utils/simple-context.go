package utils

import (
	"os"
	"os/signal"
)

type GracefulShutdown struct {
	Signals []os.Signal
}

func (gc *GracefulShutdown) Listen(done chan bool) {
	stop := make(chan os.Signal)
	signal.Notify(stop, gc.Signals...)

	for {
		select {
		case <-stop:
			done <- true
			if <-done == false {
				os.Exit(0)
			}
		}
	}
}

func NewGracefulShutDownListener(signals []os.Signal) *GracefulShutdown {
	return &GracefulShutdown{
		Signals: signals,
	}
}
