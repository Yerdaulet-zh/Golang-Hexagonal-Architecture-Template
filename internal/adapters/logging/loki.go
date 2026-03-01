package logging

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/core/ports"
	"go.opentelemetry.io/otel/trace"
)

type lokiAdapter struct {
	labels map[string]string
	url    string
}

func NewLokiLogger(url string, labels map[string]string) ports.Logger {
	return &lokiAdapter{
		url:    url,
		labels: labels,
	}
}

func (l *lokiAdapter) send(ctx context.Context, level, msg string, args ...any) {
	// Extract TraceID and SpanID from context
	spanContext := trace.SpanFromContext(ctx).SpanContext()

	var traceID, spanID string
	if spanContext.HasTraceID() {
		traceID = spanContext.TraceID().String()
	}
	if spanContext.HasSpanID() {
		spanID = spanContext.SpanID().String()
	}

	// Start the log line with Trace/Span IDs if they exist
	line := fmt.Sprintf("level=%s msg=%q", level, msg)
	if traceID != "" {
		line = fmt.Sprintf("%s trace_id=%s span_id=%s", line, traceID, spanID)
	}

	// Loop through args to create individual rows/keys
	// We assume args come in pairs: ["key", value, "key2", value2]
	for i := 0; i < len(args); i += 2 {
		if i+1 < len(args) {
			key := fmt.Sprintf("%v", args[i])
			val := fmt.Sprintf("%v", args[i+1])
			// Append each as its own key=value pair
			line = fmt.Sprintf("%s %s=%q", line, key, val)
		} else {
			// If there's an odd number of args, just append the last one
			line = fmt.Sprintf("%s extra=%q", line, args[i])
		}
	}

	ts := fmt.Sprintf("%d", time.Now().UnixNano())
	lokiMsg := map[string]any{
		"streams": []map[string]any{
			{
				"stream": l.labels,
				"values": [][]string{{ts, line}},
			},
		},
	}

	body, err := json.Marshal(lokiMsg)
	if err != nil {
		fmt.Printf("logging error: could not marshal loki msg: %v\n", err)
		return
	}

	resp, err := http.Post(l.url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		fmt.Printf("logging error: failed to send to loki: %v\n", err)
		return
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Printf("Error while closing the loki post response body: %v\n", err)
		}
	}()

	if resp.StatusCode >= 400 {
		fmt.Printf("logging error: loki returned status %d\n", resp.StatusCode)
	}
}

func (l *lokiAdapter) Debug(ctx context.Context, msg string, args ...any) {
	l.send(ctx, "Debug", msg, args...)
}

func (l *lokiAdapter) Info(ctx context.Context, msg string, args ...any) {
	l.send(ctx, "Info", msg, args...)
}

func (l *lokiAdapter) Warn(ctx context.Context, msg string, args ...any) {
	l.send(ctx, "Warn", msg, args...)
}

func (l *lokiAdapter) Error(ctx context.Context, msg string, args ...any) {
	l.send(ctx, "Error", msg, args...)
}
