package configuration

import (
	"github.com/spf13/viper"
	"fmt"
	"path/filepath"
	"strings"
)

func FileConfiguration(config, configType string) error {
	abs, err := filepath.Abs(config)
	if err != nil {
		return fmt.Errorf("error reading filepath: (%s)", err)
	}

	// get the config name
	base := filepath.Base(abs)

	// get the path
	path := filepath.Dir(abs)
	viper.SetConfigType(configType)
	viper.SetConfigName(strings.Split(base, ".")[0])
	viper.AddConfigPath(path)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Find and read the config file; Handle errors reading the config file
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading configuration file: (%s)", err)
	}

	return nil
}