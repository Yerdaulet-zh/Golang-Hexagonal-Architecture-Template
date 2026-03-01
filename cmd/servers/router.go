// Package servers provides HTTP or gRPC compatible servers to serve client requests.
package servers

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/adapters/cache/redis"
	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/adapters/handlers"
	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/adapters/handlers/middleware"
	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/core/ports"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/sdk/trace"
)

// Defining middleware
type mw func(http.Handler) http.Handler

func applyMiddlewares(h http.Handler, mws ...mw) http.Handler {
	for _, mw := range mws {
		h = mw(h)
	}
	return h
}

// MapManagementRoutes maps management-related routes (health checks, metrics) to their handlers.
func MapManagementRoutes(logger ports.Logger, client ports.Database) http.Handler {
	mux := http.NewServeMux()

	healthHdl := handlers.NewHealthHandler(client)
	mux.HandleFunc("GET /healthz", healthHdl.Healthz)
	mux.HandleFunc("GET /ready", healthHdl.Ready)

	mux.Handle("GET /metrics", promhttp.Handler())
	return mux
}

func MapBusinessRoutes(logger ports.Logger, tracer *trace.TracerProvider, rdb ports.Redis, NotificationService ports.NotificationUseCase) http.Handler {
	mux := http.NewServeMux()

	notification := handlers.NewNotificationHandler(NotificationService, logger)
	mux.HandleFunc("POST /v1/notification/email", notification.EmailNotification)

	// Middlewares
	rateLimiter := redis.NewRateLimiter(rdb)
	middlewares := []mw{
		middleware.NewIPRateLimiter(logger, rateLimiter, 100*time.Second, 1),
	}
	handler := applyMiddlewares(mux, middlewares...)
	return otelhttp.NewHandler(handler, "business-api")
}
