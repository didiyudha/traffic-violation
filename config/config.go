package config

import (
	"github.com/didiyudha/traffic-violation/internal/platform/database"
	"github.com/spf13/viper"
	"github.com/sirupsen/logrus"
)

var Configuration = new(Config)

const (
	configFilename = "config"
	configType     = "yml"
	configPath     = "."
)

type Config struct {
	DB database.Config `mapstructure:"db"`
}

func Load() error {
	viper.SetConfigName(configFilename)
	viper.AddConfigPath(configPath)
	viper.AutomaticEnv()
	viper.SetConfigType(configType)
	if err := viper.ReadInConfig(); err != nil {
		logrus.WithError(err).Info("read configuration")
		return err
	}
	err := viper.Unmarshal(Configuration)
	if err != nil {
		logrus.WithError(err).Info("unmarshal configuration")
		return err
	}
	return nil
}