// Package main is the entrypoint for the project.
// It initializes the core services and starts the gRPC, HTTP runtime.
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/cmd/servers"
	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/adapters/logging"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	if err := run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	logger := logging.NewStdoutLogger()
	routes := servers.MapManagementRoutes(logger)
	return servers.Run(ctx, logger, routes, ":2112", "Management")
}
