package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	MongoDB MongoDB
}

type MongoDB struct {
	Host     string
	Port     uint64
	User     string
	Password string
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
