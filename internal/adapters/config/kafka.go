package config

import (
	"time"

	"github.com/spf13/viper"
)

//nolint:govet
type KafkaConfig struct {
	Brokers []string
	Topic   string

	// Compression: "snappy", "gzip", "lz4", or "zstd"
	Compression string

	// Performance / Batching settings
	BatchBytes      int64         // Default: 1MB
	BatchSize       int           // Default: 100
	BatchTimeout    time.Duration // Default: 1s
	WriteBackoffMin time.Duration // Default: 100ms

	// Reliability settings
	MaxAttempts  int           // Default: 10
	RequiredAcks int           // -1 = All replicas
	WriteTimeout time.Duration // Default: 10s
	ReadTimeout  time.Duration // Default: 10s
}

func NewKafkaConfig() *KafkaConfig {
	return &KafkaConfig{
		Brokers: viper.GetStringSlice("broker.kafka.brokers"),
		Topic:   viper.GetString("broker.kafla.topic"),

		MaxAttempts:  viper.GetInt("broker.kafka.reliability.maxAttempts"),
		RequiredAcks: viper.GetInt("broker.kafka.reliability.requiredAcks"),
		WriteTimeout: viper.GetDuration("broker.kafka.reliability.writeTimeout"),
		ReadTimeout:  viper.GetDuration("broker.kafka.reliability.readTimeout"),

		BatchSize:       viper.GetInt("broker.kafka.performance.batchSize"),
		BatchBytes:      viper.GetInt64("broker.kafka.performance.batchBytes"),
		BatchTimeout:    viper.GetDuration("broker.kafka.performance.batchTimeout"),
		WriteBackoffMin: viper.GetDuration("broker.kafka.performance.writeBackoffMin"),
		Compression:     viper.GetString("broker.kafka.performance.compression"),
	}
}
