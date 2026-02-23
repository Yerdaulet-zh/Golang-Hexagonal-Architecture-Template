// Package servers provides HTTP or gRPC compatible servers to serve client requests.
package servers

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/adapters/handlers"
	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/core/ports"
)

// MapManagementRoutes maps management-related routes (health checks, metrics) to their handlers.
func MapManagementRoutes(logger ports.Logger) http.Handler {
	mux := http.NewServeMux()

	healthHdl := handlers.NewHealthHandler()
	mux.HandleFunc("GET /healthz", healthHdl.Healthz)

	mux.Handle("GET /metrics", promhttp.Handler())
	return mux
}
