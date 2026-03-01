// Package redis provides the implementation for caching and data storage
// using the Redis key-value store. It acts as a driven adapter in the
// hexagonal architecture, satisfying infrastructure ports defined by the core.
package redis

import (
	"context"

	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/adapters/config"
	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/core/domain"
	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/core/ports"
)

type redisAdapter struct {
	client *redis.Client
}

func NewRedisClient(logger ports.Logger, cfg *config.RedisConfig) ports.Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr(),
		Password: cfg.RedisPassword,
		DB:       cfg.DB,
	})
	if err := redisotel.InstrumentTracing(client); err != nil {
		logger.Error(context.TODO(), domain.LogLevelCache, "Error while enabling tracing:", err.Error())
	}
	return &redisAdapter{client: client}
}

func (a *redisAdapter) Pipeline() redis.Pipeliner {
	return a.client.Pipeline()
}

func (a *redisAdapter) Close() error {
	return a.client.Close()
}

func (a *redisAdapter) Eval(ctx context.Context, script string, keys []string, args ...interface{}) *redis.Cmd {
	return a.client.Eval(ctx, script, keys, args...)
}

func (a *redisAdapter) EvalSha(ctx context.Context, sha1 string, keys []string, args ...interface{}) *redis.Cmd {
	return a.client.EvalSha(ctx, sha1, keys, args...)
}

func (a *redisAdapter) EvalRO(ctx context.Context, script string, keys []string, args ...interface{}) *redis.Cmd {
	return a.client.EvalRO(ctx, script, keys, args...)
}
func (a *redisAdapter) EvalShaRO(ctx context.Context, sha1 string, keys []string, args ...interface{}) *redis.Cmd {
	return a.client.EvalShaRO(ctx, sha1, keys, args...)
}

func (a *redisAdapter) ScriptExists(ctx context.Context, hashes ...string) *redis.BoolSliceCmd {
	return a.client.ScriptExists(ctx, hashes...)
}

func (a *redisAdapter) ScriptLoad(ctx context.Context, script string) *redis.StringCmd {
	return a.client.ScriptLoad(ctx, script)
}
