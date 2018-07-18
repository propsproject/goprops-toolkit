package service

import (
	"github.com/propsproject/goprops-toolkit/utils/sharedconf"
	"sync"
)

type Service interface {
	Start(chan sharedconf.ConsulRegistration)
	ShutDownSig() chan *sync.WaitGroup
}

