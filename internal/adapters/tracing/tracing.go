// Package tracing provides utilities for initializing and managing OpenTelemetry
// distributed tracing within the application.
//
// It simplifies the setup of the OTLP gRPC exporter, resource attribution,
// and global tracer provider registration.
package tracing

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
)

func InitTracer() (*sdktrace.TracerProvider, error) {
	ctx := context.Background()

	// Configure the exporter (Tempo usually listens on 4317)
	exporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	// Define the service name
	res, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("golang-app"),
		),
	)

	// Create the Provider
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
	)

	// Set it globally
	otel.SetTracerProvider(tp)
	return tp, nil
}
