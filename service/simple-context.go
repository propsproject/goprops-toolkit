package utils

import (
	"os"
	"os/signal"
	"github.com/propsproject/goprops-toolkit/service"
)

type GracefulShutdown struct {
	Signals []os.Signal
}

func (gc *GracefulShutdown) Listen(doneChs []chan bool) {
	stop := make(chan os.Signal)
	signal.Notify(stop, gc.Signals...)

	for {
		select {
		case <-stop:
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

func (gc *GracefulShutdown) ListenForServices(services []service.Service) {
	chs := make([]chan bool, len(services))
	for _, service := range services {
		chs = append(chs, service.ShutDownSig())
	}

	gc.Listen(chs)
}

func NewGracefulShutDownListener(signals []os.Signal) *GracefulShutdown {
	return &GracefulShutdown{
		Signals: signals,
	}
}
