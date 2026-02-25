package config

import (
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

func NewRedisCondfig() *RedisConfig {
	return &RedisConfig{
		Host:     viper.GetString("cache.redis.host"),
		Port:     viper.GetInt("cache.redis.port"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       viper.GetInt("cache.redis.db"),
	}
}

func (cfg *RedisConfig) Addr() string {
	return cfg.Host + ":" + strconv.Itoa(cfg.Port)
}
