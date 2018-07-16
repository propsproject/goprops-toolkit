package service

import (
	"os"
	"os/signal"
)

type GracefulShutdown struct {
	Signals []os.Signal
	Stop chan os.Signal
}

func (gc *GracefulShutdown) Listen(doneChs []chan bool) {
	signal.Notify(gc.Stop, gc.Signals...)
	for {
		select {
		case <-gc.Stop:
			for _, done := range doneChs {
				done <- true
				if <-done == false {
					continue
				}
			}
			os.Exit(0)
		}
	}
}

func (gc *GracefulShutdown) ListenForServices(services []Service) {
	chs := make([]chan bool, len(services)-1)
	for _, service := range services {
		chs = append(chs, service.ShutDownSig())
	}
	gc.Listen(chs)
}

func NewGracefulShutDownListener(signals []os.Signal) *GracefulShutdown {
	return &GracefulShutdown{
		Signals: signals,
		Stop:  make(chan os.Signal),
	}
}
