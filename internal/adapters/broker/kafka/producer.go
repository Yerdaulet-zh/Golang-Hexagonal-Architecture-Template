// Package kafka provides a Kafka-based implementation of the broker port.
// It uses the segmentio/kafka-go library to handle high-performance,
// pure-Go message production and connects to the system's core logging ports.
package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/adapters/config"
	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/core/ports"
)

// Producer wraps the kafka.Writer and implements the ports.Broker interface.
type Producer struct {
	writer *kafka.Writer
}

// NewProducer initializes a production-ready Kafka writer using the provided configuration.
// It bridges the segmentio logger to the application's core logger port.
func NewProducer(cfg config.KafkaConfig, logger ports.Logger) *Producer {
	w := &kafka.Writer{
		Addr:            kafka.TCP(cfg.Brokers...),
		Topic:           cfg.Topic,
		MaxAttempts:     cfg.MaxAttempts,
		WriteBackoffMin: cfg.WriteBackoffMin,
		BatchSize:       cfg.BatchSize,
		BatchBytes:      cfg.BatchBytes,
		BatchTimeout:    cfg.BatchTimeout,
		ReadTimeout:     cfg.ReadTimeout,
		WriteTimeout:    cfg.WriteTimeout,
		RequiredAcks:    kafka.RequiredAcks(cfg.RequiredAcks),

		// Hash balancer ensures messages with the same key go to the same partition,
		// preserving order for specific entities (e.g., specific UserID).
		Balancer: &kafka.Hash{},

		// Production Loggers: bridged to the Loki/Stdout ports
		Logger:      &writerLogger{core: logger, isError: false},
		ErrorLogger: &writerLogger{core: logger, isError: true},
	}

	// Set compression based on config string to optimize bandwidth
	switch cfg.Compression {
	case "snappy":
		w.Compression = kafka.Snappy
	case "lz4":
		w.Compression = kafka.Lz4
	case "zstd":
		w.Compression = kafka.Zstd
	case "gzip":
		w.Compression = kafka.Gzip
	default:
		// Default to No Compression
	}

	return &Producer{writer: w}
}

// Publish sends a message to the configured Kafka topic.
func (p *Producer) Publish(ctx context.Context, key, msg []byte) error {
	return p.writer.WriteMessages(ctx, kafka.Message{
		Key:   key,
		Value: msg,
	})
}

// Close flushes any buffered messages and closes the Kafka connection.
func (p *Producer) Close() error {
	return p.writer.Close()
}
