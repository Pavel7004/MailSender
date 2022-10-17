package config

import (
	"time"

	"github.com/spf13/viper"
)

type MongoCfg struct {
	Uri     string        `mapstructure:"mongo_uri"`
	Timeout time.Duration `mapstructure:"mongo_timeout"`
}

type Config struct {
	Mongo MongoCfg `mapstructure:",squash"`
}

func Get() (*Config, error) {
	config := new(Config)

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.config/MailSender")
	viper.AddConfigPath("$HOME/.MailSender")
	viper.AddConfigPath(".")

	viper.SetDefault("mongo_uri", "mongodb://localhost:27017")
	viper.SetDefault("mongo_timeout", "10s")

	viper.AutomaticEnv()

	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}

	return config, nil
}

func WriteToFile(filename string) error {
	return viper.WriteConfigAs(filename)
}
