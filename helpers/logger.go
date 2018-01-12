package helpers

import "github.com/sirupsen/logrus"

type Logger interface {
	Info(string)
	Error(string)
}

type logger struct{}

func (h logger) Info(msg string) {
	logrus.Info(msg)
}

func (h logger) Error(msg string) {
	logrus.Error(msg)
}

func NewLogger() logger {
	formatter := &logrus.TextFormatter{
		FullTimestamp: true,
	}
	logrus.SetFormatter(formatter)

	return logger{}
}
