// This is designed to run http servers via net/http library
package servers

import (
	"context"
	"errors"
	"net/http"
	"time"

	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/core/ports"
)

// The function is designed to handle the bussiness and management HTTP servers
func Run(ctx context.Context, logger ports.Logger, handler http.Handler, addr string, serverName string) error {
	s := &http.Server{
		Addr:           addr,
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		logger.Info("Starting HTTP "+serverName+" server", "address", s.Addr)
		if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("HTTP "+serverName+" server failed", "error", err)
		}
	}()

	<-ctx.Done()
	logger.Info("Shutting down HTTP " + serverName + " server...")

	// Give the server 5 seconds to finish processing existing requests
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return s.Shutdown(shutdownCtx)
}
