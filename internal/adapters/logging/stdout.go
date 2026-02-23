package logging

import (
	"log/slog"
	"os"

	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/core/ports"
)

type stdoutAdapter struct {
	logger *slog.Logger
}

func NewStdoutLogger() ports.Logger {
	handler := slog.NewJSONHandler(os.Stdout, nil)
	return &stdoutAdapter{
		logger: slog.New(handler),
	}
}

func (l *stdoutAdapter) Debug(msg string, args ...any) {
	l.logger.Debug(msg, args...)
}

func (l *stdoutAdapter) Info(msg string, args ...any) {
	l.logger.Info(msg, args...)
}

func (l *stdoutAdapter) Warn(msg string, args ...any) {
	l.logger.Warn(msg, args...)
}

func (l *stdoutAdapter) Error(msg string, args ...any) {
	l.logger.Error(msg, args...)
}
