package service

import (
	"github.com/propsproject/goprops-toolkit/utils/sharedconf"
)

type Service interface {
	Start(chan sharedconf.ConsulRegistration)
	ShutDownSig() chan bool
}

