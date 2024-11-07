package main

import (
	"flag"
	"github.com/rogue0026/door-locker/internal/application"
	"github.com/rogue0026/door-locker/internal/config"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

const (
	envDev  string = "dev"
	envProd string = "prod"
)

var (
	cfgPath string
)

func main() {
	flag.StringVar(&cfgPath, "cfg", "", "path to application config")
	flag.Parse()

	appLogger := setupLogger(envDev, os.Stdout)
	appLogger.Debugf("logger initialized")
	appConfig := config.MustLoad(cfgPath)

	application.Run(appConfig, appLogger)
}

func setupLogger(env string, logsOut io.Writer) *logrus.Logger {
	var logger *logrus.Logger
	switch env {
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
