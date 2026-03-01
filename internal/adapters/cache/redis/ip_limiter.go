package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/core/ports"
)

type redisRateLimiter struct {
	client ports.Redis
}

func NewRateLimiter(client ports.Redis) ports.RateLimiter {
	return &redisRateLimiter{
		client: client,
	}
}

var rateLimitScript *redis.Script = redis.NewScript(`
	redis.call('ZREMRANGEBYSCORE', KEYS[1], 0, ARGV[1])
	local current_count = redis.call('ZCARD', KEYS[1])
	if current_count < tonumber(ARGV[4]) then
		redis.call('ZADD', KEYS[1], ARGV[2], ARGV[5])
		redis.call('EXPIRE', KEYS[1], ARGV[3])
		return 1
	end
	return 0
`)

func (r *redisRateLimiter) Allow(ctx context.Context, key string, window time.Duration, limit int64) (bool, error) {
	key = "ip_rate_limit:" + key

	now := time.Now().UnixMilli()
	windowStart := now - window.Milliseconds()
	member := fmt.Sprintf("%d", windowStart)

	result, err := rateLimitScript.Run(ctx, r.client, []string{key},
		windowStart,           // ARGV[1]
		now,                   // ARGV[2]
		int(window.Seconds()), // ARGV[3]
		limit,                 // ARGV[4]
		member,                // ARGV[5]
	).Int()
	if err != nil {
		return true, err
	}

	// result: 1 allow, 0 deny
	return result == 1, nil
}
