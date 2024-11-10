package application

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/rogue0026/door-locker/internal/config"
	"github.com/rogue0026/door-locker/internal/storage/postgres"
	"github.com/rogue0026/door-locker/internal/transport/http/handlers"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"time"
)

const (
	envDev  string = "dev"
	envProd string = "prod"
)

type BackendApplication struct {
	AppLogger  *logrus.Logger
	AppStorage *postgres.Storage
}

func New(cfg config.AppConfig) BackendApplication {
	appStorage, err := postgres.New(context.Background(), cfg.DSN)
	if err != nil {
		panic(err)
	}
	a := BackendApplication{
		AppLogger:  setupLogger(envDev, os.Stdout),
		AppStorage: &appStorage,
	}
	return a
}

func (a *BackendApplication) Run(addr string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	a.AppLogger.Debug("checking connection with database")
	err := a.AppStorage.Ping(ctx)
	if err != nil {
		panic(err)
	}
	a.AppLogger.Debug("database connection established")
	a.AppLogger.Debug("configuring http server")

	appRouter := chi.NewRouter()
	appRouter.Method(http.MethodGet, "/api/door-locks", handlers.DoorLockByLimitOffsetHandler(a.AppLogger, a.AppStorage))
	appRouter.Method(http.MethodPost, "/api/door-locks", handlers.AddDoorLockHandler(a.AppLogger, a.AppStorage))

	server := http.Server{
		Handler: appRouter,
		Addr:    addr,
	}
	a.AppLogger.Debug("configuring http server complete")
	a.AppLogger.Infof("starting listening at %s", addr)
	return server.ListenAndServe()
}

func setupLogger(env string, logsOut io.Writer) *logrus.Logger {
	var logger *logrus.Logger
	switch env {
	case envDev:
		logger = &logrus.Logger{
			Out: logsOut,
			Formatter: &logrus.TextFormatter{
				TimestampFormat: "02.01.2006 15:04:05",
				FullTimestamp:   true,
			},
			ReportCaller: true,
			Level:        logrus.DebugLevel,
		}
	case envProd:
		logger = &logrus.Logger{
			Out: logsOut,
			Formatter: &logrus.JSONFormatter{
				TimestampFormat: "02.01.2006 15:04:05",
				PrettyPrint:     false,
			},
			ReportCaller: true,
			Level:        logrus.InfoLevel,
		}
	}
	if logger == nil {
		panic("logger not initialized")
	}
	return logger
}
