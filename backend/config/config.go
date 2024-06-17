package config

import (
	"github.com/spf13/viper"
	"strings"
)

var conf *Config

type Config struct {
	Env      string   `yaml:"env"`
	MongoDB  MongoDB  `yaml:"mongodb"`
	Http     Http     `yaml:"http"`
	Telegram Telegram `yaml:"telegram"`
}

type Http struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type MongoDB struct {
	Host     string `yaml:"host"`
	Port     uint64 `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type Telegram struct {
	BotToken string `yaml:"botToken"`
}

func Init() {
	v := viper.New()
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))

	_ = v.ReadInConfig()

	var c Config
	if err := v.Unmarshal(&c); err != nil {
		panic(err)
	}

	conf = &c
}

func C() *Config {
	return conf
}
