package helpers

import (
	"github.com/sirupsen/logrus"
	"os"
	"fmt"
	"strings"
)

type Logging interface {
	Info(string)
	Error(string)
}

type Logger struct{}
type PrintlnLogger struct{}

func (PrintlnLogger) Info(msg string) {
	fmt.Println(msg)
}

func (PrintlnLogger) Error(msg string) {
	fmt.Println(msg)
}

func (h Logger) Write(p []byte) (n int, err error) {
	str := string(p)
	for _, line := range strings.Split(str, "\n") {
		logrus.Info(line)
	}
	return len(p), nil
}

func (h Logger) Info(msg string) {
	logrus.Info(msg)
}

func (h Logger) Error(msg string) {
	logrus.Error(msg)
}

func (h Logger) ForwardStdoutToStderr() {
	logrus.SetOutput(os.Stderr)
}

func NewLogger() Logger {
	formatter := &logrus.TextFormatter{
		FullTimestamp: true,
	}
	logrus.SetFormatter(formatter)

	return Logger{}
}

func NewPrinlnLogger() Logging {
	return PrintlnLogger{}
}
