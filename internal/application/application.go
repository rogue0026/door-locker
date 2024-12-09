package application

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rogue0026/door-locker/internal/config"
	pgAccounts "github.com/rogue0026/door-locker/internal/storage/accounts/postgres"
	pgLocks "github.com/rogue0026/door-locker/internal/storage/locks/postgres"
	"github.com/rogue0026/door-locker/internal/transport/http/handlers/accounts"
	"github.com/rogue0026/door-locker/internal/transport/http/handlers/locks"
	"github.com/rogue0026/door-locker/internal/transport/http/middleware"
	"github.com/rogue0026/door-locker/pkg/logging"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
)

const (
	envDev  string = "development"
	envProd string = "production"
)

type Application struct {
	Logger     *logrus.Logger
	HTTPServer *http.Server
	Accounts   *pgAccounts.Repository
	Locks      *pgLocks.Repository
	dbConnPool *pgxpool.Pool
}

func New(cfg config.AppConfig, connPool *pgxpool.Pool) (Application, error) {
	logger := logging.SetupLogger(cfg.AppEnvironment, os.Stdout)
	accountsStorage := pgAccounts.New(connPool)
	locksStorage := pgLocks.New(connPool)

	router := chi.NewRouter()
	router.Use(middleware.LoggingMiddleware(logger), middleware.AccessControl)
	router.Method(http.MethodGet, "/api/door-locks", locks.Paginated(logger, locksStorage))
	router.Method(http.MethodGet, "/api/door-locks/popular", locks.Popular(logger, locksStorage))
	router.Method(http.MethodGet, "/api/door-locks/categories", locks.Categories(logger, locksStorage))
	router.Method(http.MethodPost, "/api/door-locks", locks.Create(logger, locksStorage))
	router.Method(http.MethodDelete, "/api/door-locks/{PartNumber}", locks.Delete(logger, locksStorage))
	router.Method(http.MethodPost, "/api/accounts", accounts.Create(logger, accountsStorage))
	router.Method(http.MethodDelete, "/api/accounts", accounts.Delete(logger, accountsStorage))

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.HTTPServerHost, cfg.HTTPServerPort),
		Handler: router,
	}

	app := Application{
		Logger:     logger,
		HTTPServer: server,
		Accounts:   &accountsStorage,
		Locks:      &locksStorage,
		dbConnPool: connPool,
	}
	return app, nil
}

func (a Application) Run() error {
	return a.HTTPServer.ListenAndServe()
}

func (a Application) CloseDatabaseConnection() {
	a.dbConnPool.Close()
}

func SetupLogger(appEnvironment string, logsOut io.Writer) *logrus.Logger {
	var logger *logrus.Logger
	switch appEnvironment {
	case envDev:
		logger = &logrus.Logger{
			Out: logsOut,
			Formatter: &logrus.JSONFormatter{
				TimestampFormat: "02.01.2006 15:04:05",
				PrettyPrint:     true,
			},
			ReportCaller: true,
			Level:        logrus.DebugLevel,
		}
	case envProd:
		logger = &logrus.Logger{
			Out: logsOut,
			Formatter: &logrus.JSONFormatter{
				TimestampFormat: "02.01.2006 15:04:05",
				PrettyPrint:     true,
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

func NewConnectionPool(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	const fn = "internal.storage.postgres.New"

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return pool, nil
}
