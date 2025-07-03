package logger

import (
	"fmt"
	"time"
)

type Logger struct {
	serviceName string
}

func New(serviceName string) *Logger {
	return &Logger{serviceName: serviceName}
}

func (l *Logger) Info(message string) {
	fmt.Printf("[%s][%s][INFO] %s\n",
		time.Now().Format("2006-01-02 15:04:05"),
		l.serviceName,
		message)
}

func (l *Logger) Error(err error) {
	fmt.Printf("[%s][%s][ERROR] %v\n",
		time.Now().Format("2006-01-02 15:04:05"),
		l.serviceName,
		err)
}
