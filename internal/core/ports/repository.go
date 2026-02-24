package ports

import "context"

// Database defines the set of methods required for infrastructure
// health checks and connection lifecycle management.
// This interface is used for mocking the postgre.Client
type Database interface {
	Ping(ctx context.Context) error
	Close() error
}
