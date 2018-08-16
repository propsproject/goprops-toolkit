package service

import (
	"github.com/propsproject/goprops-toolkit/utils/sharedconf"
	"fmt"
	"os"
	"github.com/propsproject/goprops-toolkit/logging"
	"sync"
)

type RegisteredService struct {
	ID string `json:"id"`
	Registration sharedconf.ConsulRegistration
}

func (r *RegisteredService) String() string {
	return fmt.Sprintf("Name: %s, Port: %v, Consul ID: %s", r.Registration.Name, r.Registration.Port, r.ID)
}

type RegisteredServices []RegisteredService

func (r *RegisteredServices) String() string  {
	str := "Registered Services \n"
	for _, s := range *r {
		str = fmt.Sprintf("%s%s\n", str, s.String())
	}

	return str
}

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
	ShutDownSig chan bool
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
	m.Consul().Clean()
	os.Exit(0)
}

func (m *MicroService) registerServices() {
	for reg := range m.consulRegCh {
		if m.IsDevelopment {
			m.Logger().Warn(fmt.Sprintf("(%s) service running in development mode. No Consul configuration was loaded. ", m.Name))
		} else {
			_, err := m.Consul().Register(m.Name, reg, m.IsDevelopment)
			if err != nil {
				m.Logger().Warn(fmt.Sprintf("error registering service with consul: %v", err))
			}
		}
	}
}

func (m *MicroService) LoadConfigs() {
	m.once.Do(func() {
		err := m.common.LoadConfig(m.Name, &m.Config)
			if err != nil {
			m.Logger().Fatal(err)
		}
	})
}

func (m *MicroService) Logger() *logging.PLogger {
	return m.common.Logger()
}

func (m *MicroService) Consul() *sharedconf.ConsulClient  {
	return m.common.Consul()
}

func NewMicroService(name, description, version string) *MicroService {
	return &MicroService{
		consulRegCh: make(chan sharedconf.ConsulRegistration),
		Name: name,
		Description: description,
		Version: version,
		ShutDownSig: make(chan bool),
		ShutdownListener: NewGracefulShutDownListener(defaultSignals),
	}
}


