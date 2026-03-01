package logging

import (
	"context"

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

func (m *multi) Debug(ctx context.Context, msg string, args ...any) {
	for _, logger := range m.loggers {
		logger.Debug(ctx, msg, args...)
	}
}

func (m *multi) Info(ctx context.Context, msg string, args ...any) {
	for _, logger := range m.loggers {
		logger.Info(ctx, msg, args...)
	}
}

func (m *multi) Warn(ctx context.Context, msg string, args ...any) {
	for _, logger := range m.loggers {
		logger.Warn(ctx, msg, args...)
	}
}

func (m *multi) Error(ctx context.Context, msg string, args ...any) {
	for _, logger := range m.loggers {
		logger.Error(ctx, msg, args...)
	}
}
