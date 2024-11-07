package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type AppConfig struct {
	Host string `yaml:"host"`
	Port uint16 `yaml:"port"`
	DSN  string `yaml:"dsn"`
}

func MustLoad(configPath string) AppConfig {
	_, err := os.Lstat(configPath)
	if err != nil {
		panic(err)
	}
	cfg := AppConfig{}
	err = cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}
