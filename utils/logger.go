package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	Logger *logrus.Logger
}

var LoggerSingleton *Logger

func (l *Logger) Init() {
	l.Logger = logrus.New()
	l.Logger.SetOutput(os.Stdout)
	l.Logger.SetLevel(logrus.InfoLevel)
}
