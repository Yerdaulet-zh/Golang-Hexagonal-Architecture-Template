package servers

import (
	"context"
	"net/http"
	"testing"
	"time"
)

type mockLogger struct{}

func (m *mockLogger) Info(msg string, args ...any)  {}
func (m *mockLogger) Error(msg string, args ...any) {}
func (m *mockLogger) Debug(msg string, args ...any) {}
func (m *mockLogger) Warn(msg string, args ...any)  {}

// TestRun verifies that the server starts, accepts a request,
// and shuts down gracefully when the context is canceled.
func TestRun(t *testing.T) {
	// 1. Setup dependencies
	logger := &mockLogger{}
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	ctx, cancel := context.WithCancel(context.Background())

	// Use port :0 to get a random available port
	addr := "127.0.0.1:0"

	// Start the server in a goroutine
	errCh := make(chan error, 1)
	go func() {
		errCh <- Run(ctx, logger, handler, addr, "TestServer")
	}()

	// Give the server a tiny bit of time to bind to the port
	time.Sleep(100 * time.Millisecond)

	// Trigger Shutdown
	cancel()

	// Verify Shutdown
	select {
	case err := <-errCh:
		if err != nil {
			t.Errorf("Run() returned unexpected error on shutdown: %v", err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("Server timed out during shutdown")
	}
}
