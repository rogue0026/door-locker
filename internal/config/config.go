package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type AppConfig struct {
	AppEnvironment string `env:"APP_ENVIRONMENT"`
	LogLevel       string `env:"LOG_LEVEL"`
	HTTPServerHost string `env:"HTTP_SERVER_HOST"`
	HTTPServerPort uint   `env:"HTTP_SERVER_PORT"`
	DBHost         string `env:"DB_HOST"`
	DBPort         uint   `env:"DB_PORT"`
	DBUser         string `env:"DB_USER"`
	DBUserPassword string `env:"DB_USER_PASSWORD"`
	DatabaseName   string `env:"DATABASE_NAME"`
}

func MustLoad() AppConfig {
	appConfig := AppConfig{}
	err := cleanenv.ReadEnv(&appConfig)
	if err != nil {
		panic(err)
	}
	return appConfig
}
