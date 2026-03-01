package ports

import (
	"context"
	"time"
)

type RateLimiter interface {
	Allow(ctx context.Context, key string, window time.Duration, limit int64) (bool, error)
}
