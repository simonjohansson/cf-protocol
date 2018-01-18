package helpers

import (
	"os"
	"fmt"
	"strings"
)

type Logger struct {
	pipe *os.File
}

func (p Logger) Info(msg string) {
	fmt.Fprintln(p.pipe, msg)
}

func (p Logger) Error(msg string) {
	fmt.Fprintln(p.pipe, msg)
}

func (p *Logger) ForwardStdoutToStderr() {
	p.pipe = os.Stderr
}

func (h Logger) Write(p []byte) (n int, err error) {
	str := string(p)
	for _, line := range strings.Split(str, "\n") {
		h.Info(line)
	}
	return len(p), nil
}

func NewLogger() Logger {
	return Logger{
		pipe: os.Stdout,
	}
}
