package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func NewLogger() *logrus.Logger {

	var level logrus.Level
	level = LogLevel("info")
	logger := &logrus.Logger{
		Out:   os.Stdout,
		Level: level,
		Formatter: &logrus.TextFormatter{
			DisableColors: true,
			FullTimestamp: true,
		},
	}
	Logger = logger
	return Logger
}

func LogLevel(lvl string) logrus.Level {
	switch lvl {
	case "info":
		return logrus.InfoLevel
	case "error":
		return logrus.ErrorLevel
	default:
		panic("Not supported")
	}
}
