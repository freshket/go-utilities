package logger

import (
	"github.com/sirupsen/logrus"
	"go.elastic.co/ecslogrus"
)

func CreateLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&ecslogrus.Formatter{})
	return logger
}
