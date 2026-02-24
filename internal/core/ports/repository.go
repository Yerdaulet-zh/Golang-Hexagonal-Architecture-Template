package ports

import "context"

type Database interface {
	Ping(ctx context.Context) error
	Close() error
}
