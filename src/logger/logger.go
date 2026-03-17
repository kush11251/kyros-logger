package logger

import (
	"fmt"
	"log"
)

type Logger struct {
	cfg *config.Config
}

func NewLogger(cfg *config.Config) (*Logger, error) {
	return &Logger{cfg: cfg}, nil
}

func (l *Logger) Infof(format string, v ...interface{}) {
	log.Printf(format, v...)
}