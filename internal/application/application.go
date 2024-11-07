package application

import (
	"context"
	"fmt"
	"github.com/rogue0026/door-locker/internal/config"
	"github.com/rogue0026/door-locker/internal/storage/postgres"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Run(cfg config.AppConfig, appLogger *logrus.Logger) {
	_, err := postgres.New(context.Background(), cfg.DSN)
	if err != nil {
		panic(err.Error())
	}
	appLogger.Info("database connection initialized")

	_ = http.Server{
		Handler: nil,
		Addr:    fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
	}
}

func appRouter() {

}
