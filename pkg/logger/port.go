package logger

import (
	"net/http"

	"github.com/elliot14A/tcorp/assessment/pkg/config"
)

type LoggerPort interface {
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message string, args ...interface{})
	GetHTTPMiddleWare() func(next http.Handler) http.Handler
	Fatal(message string, args ...interface{})
	SetConfig(config config.LoggerConfig) error
}

func New() LoggerPort {
	return NewSlogAdapter()
}
