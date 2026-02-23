// Package logging provides implementations for system logging adapters.
package logging

import (
	"log/slog"
	"os"

	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/core/ports"
)

// stdoutAdapter is an internal implementation of the ports.Logger interface
// that writes logs to standard output in JSON format.
type stdoutAdapter struct {
	logger *slog.Logger
}

// NewStdoutLogger initializes and returns a new Logger that outputs to stdout.
// It uses slog's JSONHandler for structured logging.
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
