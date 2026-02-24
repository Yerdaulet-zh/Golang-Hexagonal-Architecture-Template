package ports

import "context"

type Broker interface {
	Publish(ctx context.Context, key []byte, message []byte) error
	Close() error
}
