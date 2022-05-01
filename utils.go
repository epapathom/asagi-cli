package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func initializeLogger() (logger logrus.Logger) {
	logger = *logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.InfoLevel)

	return
}
