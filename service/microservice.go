package service

import (
	"github.com/propsproject/goprops-toolkit/utils/sharedconf"
	"fmt"
	"os"
	"github.com/spf13/viper"
	"github.com/propsproject/goprops-toolkit/logger"
)

type MicroService struct {
	Name string
	Description string
	Version string
	common sharedconf.Config
	ShutdownListener *GracefulShutdown
	services []Service
	Config map[string]interface{}
	consulRegCh chan sharedconf.ConsulRegistration
	IsDevelopment bool
}

var defaultSignals = []os.Signal{os.Kill, os.Interrupt}

func (m *MicroService) String() string {
	return fmt.Sprintf("%s version: %s", m.Name, m.Version)
}

func (m *MicroService) AddServices(services ...Service)  {
	m.services = append(m.services, services...)
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

func (m *MicroService) initConfigs(configPath string) error {
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("Fatal error config file: %s \n", err)
	}

	err = viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("Fatal error config file: %s \n", err)
	}

	err = viper.Unmarshal(m.Config)
	if err != nil {
		return fmt.Errorf("Fatal error config file: %s \n", err)
	}

	m.Name = m.Config["name"].(string)
	m.Description = m.Config["description"].(string)
	m.Version = m.Config["version"].(string)

	err = m.common.LoadConfig(m.Name)
	if err != nil {
		return fmt.Errorf("Fatal error loading common configs: %s \n", err)
	}

	return nil
}


func (m *MicroService) Logger() *logger.Wrapper  {
	return m.common.Logger()
}

func (m *MicroService) Consul() *sharedconf.ConsulClient  {
	return m.common.Consul()
}

func NewMicroService(configPath string) (*MicroService, error) {
	microService := &MicroService{
		ShutdownListener: NewGracefulShutDownListener(defaultSignals),
		consulRegCh: make(chan sharedconf.ConsulRegistration),
	}

	var err error
	if configPath !=  "" {
		err = microService.initConfigs(configPath)
	} else {
		microService.IsDevelopment = true
	}

	return microService, err
}


