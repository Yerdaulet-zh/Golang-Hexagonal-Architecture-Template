// Package logging provides implementations for system logging adapters.
package logging

import (
	"context"
	"log/slog"
	"os"

	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/core/ports"
	"go.opentelemetry.io/otel/trace"
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

func (l *stdoutAdapter) appendTraceInfo(ctx context.Context, args []any) []any {
	spanContext := trace.SpanFromContext(ctx).SpanContext()
	if spanContext.HasTraceID() {
		args = append(args,
			slog.String("trace_id", spanContext.TraceID().String()),
			slog.String("span_id", spanContext.SpanID().String()),
		)
	}
	return args
}

func (l *stdoutAdapter) Debug(ctx context.Context, msg string, args ...any) {
	args = l.appendTraceInfo(ctx, args)
	l.logger.Debug(msg, args...)
}

func (l *stdoutAdapter) Info(ctx context.Context, msg string, args ...any) {
	args = l.appendTraceInfo(ctx, args)
	l.logger.Info(msg, args...)
}

func (l *stdoutAdapter) Warn(ctx context.Context, msg string, args ...any) {
	args = l.appendTraceInfo(ctx, args)
	l.logger.Warn(msg, args...)
}

func (l *stdoutAdapter) Error(ctx context.Context, msg string, args ...any) {
	args = l.appendTraceInfo(ctx, args)
	l.logger.Error(msg, args...)
}
