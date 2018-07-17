package service

import (
	"github.com/propsproject/goprops-toolkit/utils/sharedconf"
	"fmt"
	"os"
	"github.com/spf13/viper"
	"github.com/propsproject/goprops-toolkit/logger"
	"sync"
)

type MicroService struct {
	Name string
	Description string
	Version string
	common sharedconf.Config
	ShutdownListener *GracefulShutdown
	services []Service
	Config map[string]interface{}
	ConfigPath string
	consulRegCh chan sharedconf.ConsulRegistration
	IsDevelopment bool
	once sync.Once
}

var defaultSignals = []os.Signal{os.Kill, os.Interrupt}

func (m *MicroService) String() string {
	return fmt.Sprintf("%s version: %s", m.Name, m.Version)
}

func (m *MicroService) AddServices(services ...Service) *MicroService {
	m.services = append(m.services, services...)
	return m
}

func (m *MicroService) Run() {
	for _, service := range m.services {
		go service.Start(m.consulRegCh)
	}

	go m.registerServices()
	m.ShutdownListener.ListenForServices(m.services)
}

func (m *MicroService) registerServices() {
	for reg := range m.consulRegCh {
		name := fmt.Sprintf("%s.%s", viper.GetString("name"), reg.Name)
		if m.IsDevelopment {
			m.Logger().Warn(fmt.Sprintf("(%s) service running in development mode. No Consul configuration was loaded. ", name))
		} else {
			m.Consul().Register(reg.Name, reg.Port)
		}
	}
	close(m.consulRegCh)
}

func (m *MicroService) LoadConfigs() {
	m.once.Do(func() {
		m.common.LoadConfig(m.Name)
	})
}

func (m *MicroService) Logger() *logger.Wrapper  {
	return m.common.Logger()
}

func (m *MicroService) Consul() *sharedconf.ConsulClient  {
	return m.common.Consul()
}

func NewMicroService() *MicroService {
	return &MicroService{
		ShutdownListener: NewGracefulShutDownListener(defaultSignals),
		consulRegCh: make(chan sharedconf.ConsulRegistration),
	}
}


