package sharedconf

import (
	"fmt"

	"github.com/caarlos0/env"
	propsLogger "github.com/propsproject/goprops-toolkit/logger"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"sync"
)

const (
	configType    = "json"
	viperProvider = "consul"

	microservicesKey = "microservices"
)

type Config struct {
	consulI
	loggerI
	IsProduction bool `env:"false"`
}

type consulI struct {
	instance *ConsulClient
	Address  string `env:"CONSUL_URI" envDefault:"127.0.0.1:8500"`
	once     sync.Once
}

type loggerI struct {
	instance *propsLogger.Wrapper
	once     sync.Once
}

func (c *Config) logger() *propsLogger.Wrapper {
	c.loggerI.once.Do(func() {
		c.loggerI.instance = propsLogger.NewLogger()
	})
	return c.loggerI.instance
}

func (c *Config) consulClient() *ConsulClient {
	c.consulI.once.Do(func() {
		client, err := newConsulClient(c.consulI.Address)
		if err != nil {
			c.Logger().Fatal(err)
		}
		c.consulI.instance = client
	})
	return c.consulI.instance
}

func (c *Config) LoadConfig(microService string) error {
	err := env.Parse(&c.consulI)
	if err != nil {
		return err
	}

	err = viper.AddRemoteProvider(viperProvider, c.consulI.Address, fmt.Sprintf("%s/%s", microservicesKey, microService))
	if err != nil {
		return err
	}

	viper.SetConfigType(configType)
	err = viper.ReadRemoteConfig()
	if err != nil {
		return err
	}

	c.logger().Info(fmt.Sprintf("%s %s %s", "Config loaded for", microService, "from consul"))
	c.consulClient()
	return nil
}

func (c *Config) Logger() *propsLogger.Wrapper {
	return c.logger()
}

func (c *Config) Consul() *ConsulClient {
	return c.consulClient()
}
