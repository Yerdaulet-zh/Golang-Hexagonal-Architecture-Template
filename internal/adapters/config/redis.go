package config

import (
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type RedisConfig struct {
	RedisPassword string
	Host          string
	Port          int
	DB            int
}

func NewRedisCondfig() *RedisConfig {
	return &RedisConfig{
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
		Host:          viper.GetString("cache.redis.host"),
		Port:          viper.GetInt("cache.redis.port"),
		DB:            viper.GetInt("cache.redis.db"),
	}
}

func (cfg *RedisConfig) Addr() string {
	return cfg.Host + ":" + strconv.Itoa(cfg.Port)
}
