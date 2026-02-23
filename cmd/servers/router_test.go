package servers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/adapters/logging"
)

func TestMapManagementRoutes(t *testing.T) {
	logger := logging.NewStdoutLogger()
	router := MapManagementRoutes(logger)

	// Create a list of test cases
	tests := []struct {
		name           string
		url            string
		expectedStatus int
	}{
		{"Health check endpoint", "/healthz", http.StatusOK},
		{"Metrics endpoint", "/metrics", http.StatusOK},
		{"Non-existent endpoint", "/not-found", http.StatusNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// httptest.NewRecorder() acts like a ResponseWriter
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", tt.url, nil)

			router.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}
		})
	}
}
