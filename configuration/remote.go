package configuration

import (
	"github.com/spf13/viper"
	"fmt"
)

// ConsulConfiguration load configuration from remote consul instance
func ConsulConfiguration(address, path, configType string, watch bool) error {
	viper.AddRemoteProvider("consul", address, path)
	viper.SetConfigType(configType)

	err := viper.ReadRemoteConfig()
	if err != nil {
		return fmt.Errorf("error reading remote cionfiguration (%s)", err)
	}

	if watch {
		err := watchRemoteConfig()
		if err != nil {
			return fmt.Errorf("error watching remote cionfiguration (%s)", err)
		}
	}

	return nil
}

func watchRemoteConfig() error {
	return viper.WatchRemoteConfig()
}
