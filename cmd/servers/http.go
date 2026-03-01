// Package servers provides functionality to run and manage HTTP servers.
package servers

import (
	"context"
	"errors"
	"net/http"
	"time"

	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/core/ports"
)

// Run handles the business and management HTTP servers, supporting graceful shutdown.
func Run(ctx context.Context, logger ports.Logger, handler http.Handler, addr string, gracefullShutdown time.Duration, serverName string) error {
	s := &http.Server{
		Addr:           addr,
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		logger.Info(ctx, "Starting HTTP "+serverName+" server", "address", s.Addr)
		if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error(ctx, "HTTP "+serverName+" server failed", "error", err)
		}
	}()

	<-ctx.Done()
	logger.Info(ctx, "Shutting down HTTP "+serverName+" server...")

	// Give the server 5 seconds to finish processing existing requests
	shutdownCtx, cancel := context.WithTimeout(context.Background(), gracefullShutdown)
	defer cancel()

	return s.Shutdown(shutdownCtx)
}
