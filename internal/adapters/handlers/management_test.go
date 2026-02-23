package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestHealthzHandler verifies that the Healthz handler returns a 200 OK
// status code and the correct "OK" response body.
func TestHealthzHandler(t *testing.T) {
	h := NewHealthHandler()

	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	w := httptest.NewRecorder()

	h.Healthz(w, req)

	resp := w.Result()
	defer func() {
		if err := resp.Body.Close(); err != nil {
			t.Error("Failed to close the respose body:", err)
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}

	if string(body) != "OK" {
		t.Errorf("expected 'OK', got %s", string(body))
	}
}
