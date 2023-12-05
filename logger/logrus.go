package logger

import "github.com/sirupsen/logrus"

type Logger interface {
	Errorf(format string, args logrus.Fields)
	Info(args logrus.Fields)
}
type loggerWrapper struct {
	lw *logrus.Logger
}

func (logger *loggerWrapper) Errorf(format string, args logrus.Fields) {
	logger.lw.WithFields(args).Errorf(format)
}

func (logger *loggerWrapper) Info(args logrus.Fields) {
	logger.lw.WithFields(args).Info()
}

func NewLogger() Logger {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	return &loggerWrapper{
		lw: logrus.New(),
	}
}
