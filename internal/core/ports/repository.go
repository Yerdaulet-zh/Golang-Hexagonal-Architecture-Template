package ports

import (
	"context"

	"gorm.io/gorm"
)

// Database defines the set of methods required for infrastructure
// health checks and connection lifecycle management.
// This interface is used for mocking the postgre.Client
type Database interface {
	Ping(ctx context.Context) error
	Close() error
	GetGormDB() *gorm.DB
}
