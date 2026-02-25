// Package redis provides the implementation for caching and data storage
// using the Redis key-value store. It acts as a driven adapter in the
// hexagonal architecture, satisfying infrastructure ports defined by the core.
package redis

import (
	"github.com/redis/go-redis/v9"
	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/adapters/config"
)

func NewRedisClient(cfg *config.RedisConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cfg.Addr(),
		Password: cfg.RedisPassword,
		DB:       cfg.DB,
	})
}
