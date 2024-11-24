package application

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/rogue0026/door-locker/internal/config"
	"github.com/rogue0026/door-locker/internal/storage/postgres"
	"github.com/rogue0026/door-locker/internal/transport/http/handlers"
	"github.com/rogue0026/door-locker/internal/transport/http/middleware"
	"github.com/sirupsen/logrus"
)

const (
	envDev  string = "development"
	envProd string = "production"
)

type BackendApplication struct {
	AppLogger  *logrus.Logger
	AppStorage *postgres.Storage
	HTTPServer *http.Server
}

func New(cfg config.AppConfig, appStorage *postgres.Storage) BackendApplication {
	appLogger := setupLogger(cfg.AppEnvironment, os.Stdout)

	loggingMw := middleware.LoggingMiddleware(appLogger)

	appRouter := chi.NewRouter()
	appRouter.Use(loggingMw, middleware.AccessControl)

	appRouter.Method(http.MethodGet, "/api/door-locks", handlers.DoorLockByLimitOffsetHandler(appLogger, appStorage))
	appRouter.Method(http.MethodGet, "/api/door-locks/popular", handlers.PopularDoorLocks(appLogger, appStorage))
	appRouter.Method(http.MethodGet, "/api/door-locks/categories", handlers.DoorLocksCategories(appLogger, appStorage))
	appRouter.Method(http.MethodPost, "/api/door-locks", handlers.AddDoorLock(appLogger, appStorage))
	appRouter.Method(http.MethodDelete, "/api/door-locks", handlers.DeleteDoorLock(appLogger, appStorage))
	appRouter.Method(http.MethodPost, "/api/accounts", handlers.RegisterAccount(appLogger, appStorage))
	appRouter.Method(http.MethodDelete, "/api/accounts", handlers.DeleteAccount(appLogger, appStorage))

	addr := fmt.Sprintf("%s:%d", cfg.HTTPServerHost, cfg.HTTPServerPort)
	server := &http.Server{
		Handler: appRouter,
		Addr:    addr,
	}
	a := BackendApplication{
		AppLogger:  setupLogger(cfg.AppEnvironment, os.Stdout),
		AppStorage: appStorage,
		HTTPServer: server,
	}
	return a
}

func (a *BackendApplication) Run() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	a.AppLogger.Debug("checking connection with database")
	err := a.AppStorage.Ping(ctx)
	if err != nil {
		panic(err)
	}
	a.AppLogger.Debug("database connection established")
	a.AppLogger.Infof("starting server at %s", a.HTTPServer.Addr)
	return a.HTTPServer.ListenAndServe()
}

func setupLogger(appEnvironment string, logsOut io.Writer) *logrus.Logger {
	var logger *logrus.Logger
	switch appEnvironment {
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
