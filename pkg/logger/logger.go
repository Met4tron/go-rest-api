package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func init() {
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}
