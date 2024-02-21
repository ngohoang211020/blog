package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

type Configuration struct {
	Port           string `mapstructure:"PORT"`
	Host           string `mapstructure:"HOST"`
	DSN            string `mapstructure:"DSN"`
	MaxOpenConn    int    `mapstructure:"MAX_OPEN_CONNECTION"`
	MaxIdleConn    int    `mapstructure:"MAX_IDLE_CONNECTION"`
	MaxIdleTimeout string `mapstructure:"MAX_IDLE_TIMEOUT"`
}

var Config Configuration

func (config *Configuration) LoadConfig() error {
	pwd, err := os.Getwd()

	if err != nil {
		log.Fatalf("failed to get current working dir. %v", err)
	}

	viper.AddConfigPath(pwd)
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(config)
	return err
}
