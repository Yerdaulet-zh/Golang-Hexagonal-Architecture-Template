// Package main is the entrypoint for the project.
// It initializes the core services and starts the gRPC, HTTP runtime.
package main

import (
	"context"
	"os/signal"
	"syscall"

	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/cmd/servers"
	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/adapters/logging"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	logger := logging.NewStdoutLogger()
	mapManagementRoutes := servers.MapManagementRoutes(logger)

	if err := servers.Run(ctx, logger, mapManagementRoutes, ":2112", "Management"); err != nil {
		logger.Error("HTTP Management server error while shutting down", "error", err)
	}
}
