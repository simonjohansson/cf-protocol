package logger

import "github.com/sirupsen/logrus"

type Logger interface {
	Info(string)
	Error(string)
}

// Impl
//
type RealLogger struct{}

func (h RealLogger) Info(msg string) {
	logrus.Info(msg)
}

func (h RealLogger) Error(msg string) {
	logrus.Error(msg)
}

func NewLogger() RealLogger {
	formatter := &logrus.TextFormatter{
		FullTimestamp: true,
	}
	logrus.SetFormatter(formatter)

	return RealLogger{}
}

// Mock
type MockLogger struct{}

func (MockLogger) Info(string) {
}

func (MockLogger) Error(string) {
}

func NewMockLogger() MockLogger {
	return MockLogger{}
}
