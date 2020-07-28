package logger

import (
	"fmt"
	"os"
)

type Logger struct{}

func New() *Logger {
	return &Logger{}
}

func (l *Logger) Printf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, format, args...)
}

func (l *Logger) Println(args ...interface{}) {
	fmt.Fprintln(os.Stdout, args...)
}

func (l *Logger) Error(err error) {
	fmt.Fprintln(os.Stderr, err)
}
