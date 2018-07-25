package sharedconf

import (
	"fmt"

	"github.com/caarlos0/env"
	propsLogger "github.com/propsproject/goprops-toolkit/logger"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"sync"
	"time"
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

type viperRuntime struct {
	runtime *viper.Viper
	address string
	kv string
	logger *propsLogger.Wrapper
	conf *map[string]interface{}
}

func (v *viperRuntime) WatchConfig() error {
	err := v.runtime.WatchRemoteConfigOnChannel()
	if err != nil {
		v.logger.Fatal(err)
	}

	for {
		time.Sleep(time.Second * 5)
		v.runtime.Unmarshal(v.conf)
	}
}

func (v *viperRuntime) LoadConfig() error {
	err := v.runtime.AddRemoteProvider(viperProvider, v.address, v.kv)
	if err != nil {
		return err
	}

	v.runtime.SetConfigType(configType)

	err = v.runtime.ReadRemoteConfig()
	if err != nil {
		return err
	}

	v.runtime.Unmarshal(v.conf)
	return err
}

func NewRuntimeViper(address, kv string, logger *propsLogger.Wrapper, runtimeConf *map[string]interface{}) *viperRuntime {
	return &viperRuntime{
		runtime: viper.New(),
		address: address,
		kv: kv,
		logger: logger,
		conf: runtimeConf,
	}
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

func (c *Config) LoadConfig(microService string, runtimeConf *map[string]interface{}) error {
	err := env.Parse(&c.consulI)
	if err != nil {
		return err
	}

	v := NewRuntimeViper(c.consulI.Address, fmt.Sprintf("%s/%s", microservicesKey, microService), c.logger(), runtimeConf)
	err = v.LoadConfig()
	if err != nil {
		return err
	}

	go v.WatchConfig()

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
