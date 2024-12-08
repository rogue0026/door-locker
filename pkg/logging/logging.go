package logging

import (
	"github.com/sirupsen/logrus"
	"io"
)

const (
	envDev  string = "development"
	envProd string = "production"
)

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
