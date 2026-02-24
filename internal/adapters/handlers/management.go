// Package handlers provides HTTP request handlers for the management server,
// including health checks and readiness probes.
package handlers

import (
	"net/http"

	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/core/ports"
)

// HealthHandler defines the dependencies for health checks
type HealthHandler struct {
	dbClient ports.Database
}

// NewHealthHandler creates a new instance of HealthHandler.
func NewHealthHandler(dbClient ports.Database) *HealthHandler {
	return &HealthHandler{dbClient: dbClient}
}

// Healthz handles the health check requests.
func (h *HealthHandler) Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	//nolint:errcheck,gosec // No need to check errors when writing simple health status
	w.Write([]byte("OK"))
}

// Ready handles the accessibility of PostgreSQL database
func (h *HealthHandler) Ready(w http.ResponseWriter, r *http.Request) {
	if err := h.dbClient.Ping(r.Context()); err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		//nolint:errcheck,gosec
		w.Write([]byte("Service Unavailable: Database unreachable"))
		return
	}
	w.WriteHeader(http.StatusOK)
	//nolint:errcheck,gosec
	w.Write([]byte("Ready"))
}
