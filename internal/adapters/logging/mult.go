package logging

import (
	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/core/ports"
)

type multi struct {
	loggers []ports.Logger
}

func NewMultiLogger(loggers ...ports.Logger) ports.Logger {
	return &multi{
		loggers: loggers,
	}
}

func (m *multi) Debug(msg string, args ...any) {
	for _, logger := range m.loggers {
		logger.Debug(msg, args...)
	}
}

func (m *multi) Info(msg string, args ...any) {
	for _, logger := range m.loggers {
		logger.Info(msg, args...)
	}
}

func (m *multi) Warn(msg string, args ...any) {
	for _, logger := range m.loggers {
		logger.Warn(msg, args...)
	}
}

func (m *multi) Error(msg string, args ...any) {
	for _, logger := range m.loggers {
		logger.Error(msg, args...)
	}
}
