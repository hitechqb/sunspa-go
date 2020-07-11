package common

import (
	"github.com/sirupsen/logrus"
	"os"
)

type Logger struct {
	*logrus.Logger
}

func NewLogger() Logger {
	var logger = logrus.New()
	logger.Out = os.Stdout
	logger.Formatter = &logrus.JSONFormatter{
		TimestampFormat: "15:04:05",
	}

	return Logger{logger}
}
