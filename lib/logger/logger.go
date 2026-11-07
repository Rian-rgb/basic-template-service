package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func SetupLogger(env string) {
	log = logrus.New()

	log.SetFormatter(&logrus.JSONFormatter{
		DisableTimestamp: false,
		PrettyPrint:      env == "development",
	})

	log.SetOutput(os.Stdout)

	if env == "production" {
		log.SetLevel(logrus.InfoLevel)
	} else {
		log.SetLevel(logrus.DebugLevel)
	}

	log.Info("Logger successfully initiated")
}

func Info(message string, args ...interface{}) {
	log.Infof(message, args...)
}

func Warn(message string, args ...interface{}) {
	log.Warnf(message, args...)
}

func Error(message string, args ...interface{}) {
	log.Errorf(message, args...)
}
