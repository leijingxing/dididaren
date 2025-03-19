package logger

import (
	"fmt"
	"time"
)

type Logger struct {
	Level string
}

func NewLogger(level string) *Logger {
	return &Logger{
		Level: level,
	}
}

func (l *Logger) Info(format string, args ...interface{}) {
	if l.Level == "debug" || l.Level == "info" {
		fmt.Printf("[INFO] %s %s\n", time.Now().Format("2006-01-02 15:04:05"), fmt.Sprintf(format, args...))
	}
}

func (l *Logger) Error(format string, args ...interface{}) {
	fmt.Printf("[ERROR] %s %s\n", time.Now().Format("2006-01-02 15:04:05"), fmt.Sprintf(format, args...))
}

func (l *Logger) Debug(format string, args ...interface{}) {
	if l.Level == "debug" {
		fmt.Printf("[DEBUG] %s %s\n", time.Now().Format("2006-01-02 15:04:05"), fmt.Sprintf(format, args...))
	}
}

func (l *Logger) Warn(format string, args ...interface{}) {
	if l.Level == "debug" || l.Level == "info" || l.Level == "warn" {
		fmt.Printf("[WARN] %s %s\n", time.Now().Format("2006-01-02 15:04:05"), fmt.Sprintf(format, args...))
	}
}
