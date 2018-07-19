package service

import (
	"os"
	"os/signal"
	"sync"
)

type GracefulShutdown struct {
	Signals []os.Signal
}

func (gc *GracefulShutdown) NotifyShutdown(done chan bool, wg *sync.WaitGroup) {
	done <- true
	for {
		select {
		case <- done:
			wg.Done()
			return
		}
	}
}

func (gc *GracefulShutdown) ListenForServices(services []Service) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)


	for {
		select {
		case <- c:
			var wg sync.WaitGroup
			go func() {
				for _, service := range services {
					wg.Add(1)
					service.ShutDownSig() <- &wg
				}
			}()
			wg.Wait()
			return
		}
	}
}

func NewGracefulShutDownListener(signals []os.Signal) *GracefulShutdown {
	return &GracefulShutdown{
		Signals: signals,
	}
}
