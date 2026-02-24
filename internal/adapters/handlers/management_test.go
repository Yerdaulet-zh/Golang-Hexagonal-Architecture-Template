package handlers

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockDB struct {
	pingErr error
}

func (m *mockDB) Ping(ctx context.Context) error {
	return m.pingErr
}

func (m *mockDB) Close() error {
	return nil
}

// TestHealthzHandler verifies that the Healthz handler returns a 200 OK
// status code and the correct "OK" response body.
func TestReadyHandler(t *testing.T) {
	tests := []struct {
		mockErr        error
		name           string
		expectedBody   string
		expectedStatus int
	}{
		{
			name:           "Success - Database is reachable",
			expectedBody:   "Ready",
			mockErr:        nil,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Failure - Database unreachable",
			expectedBody:   "Service Unavailable: Database unreachable",
			mockErr:        fmt.Errorf("connection refused"),
			expectedStatus: http.StatusServiceUnavailable,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			db := &mockDB{pingErr: tt.mockErr}
			h := NewHealthHandler(db)
			req := httptest.NewRequest(http.MethodGet, "/ready", nil)
			w := httptest.NewRecorder()

			// Act
			h.Ready(w, req)

			// Assert
			resp := w.Result()
			body, _ := io.ReadAll(resp.Body)

			if resp.StatusCode != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, resp.StatusCode)
			}

			if string(body) != tt.expectedBody {
				t.Errorf("expected body %q, got %q", tt.expectedBody, string(body))
			}
		})
	}
}
