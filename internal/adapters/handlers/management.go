// Package handlers provides HTTP request handlers for the management server,
// including health checks and readiness probes.
package handlers

import (
	"net/http"
)

// HealthHandler defines the dependencies for health checks
type HealthHandler struct{}

// NewHealthHandler creates a new instance of HealthHandler.
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// Healthz handles the health check requests.
func (h *HealthHandler) Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	//nolint:errcheck,gosec // No need to check errors when writing simple health status
	w.Write([]byte("OK"))
}
