// Package main is the entrypoint for the project.
// It initializes the core services and starts the gRPC, HTTP runtime.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/redis/go-redis/v9"
	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/cmd/servers"
	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/adapters/broker/kafka"
	redisCache "gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/adapters/cache/redis"
	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/adapters/config"
	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/adapters/logging"
	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/adapters/repository/postgre"
	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/core/ports"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	logger, client, rdb, _ := loadComponents()

	if err := run(ctx, logger, client, rdb); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, logger ports.Logger, client ports.Database, rdb *redis.Client) error {
	defer func() {
		logger.Info("Closing infrastructure connections...")
		if err := client.Close(); err != nil {
			logger.Error("Postgre close error", "error", err)
		}
		if err := rdb.Close(); err != nil {
			logger.Error("Redis close error", "error", err)
		}
		logger.Info("Done")
	}()

	logger.Info("Loading HTTP Server config")
	httpConfig := config.NewHttpConfig()
	logger.Info("Successfully loaded HTTP Server config")

	// userRepo := postgre.NewUserRepository(client.DB, &logger)
	// userService := service.NewUserService(userRepo)

	// mapBusinessHandler := servers.MapBusinessRoutes(logger, rdb, userService)
	mapManagementRoutes := servers.MapManagementRoutes(logger, client)

	go func() {
		if err := servers.Run(ctx, logger, mapManagementRoutes, httpConfig.HttpManagementAddr(), "Management"); err != nil {
			logger.Error("HTTP Management server error while shutting down", "error", err)
		}
	}()

	if err := servers.Run(ctx, logger, mapManagementRoutes, httpConfig.HttpBusinessAddr(), "Business"); err != nil { // mapBusinessHandler
		logger.Error("HTTP Business server error while shutting down", "error", err)
		os.Exit(1)
	}

	logger.Info("Application exited cleanly")
	return nil
}

func loadComponents() (ports.Logger, *postgre.Client, *redis.Client, ports.Broker) {
	// Configuration
	cfg, err := config.NewLoggingConfig()
	if err != nil {
		log.Fatalf("Error initializing config: %v", err)
	}

	// Logger
	logger := logging.NewLogger(cfg)
	logger.Info("Logging successfully configured to use the adapter: ", cfg.Adapter())

	// PostgreSQL
	logger.Info("Loading PostgreSQL config")
	postgreConfig, err := config.NewDefaultDBConfig()
	if err != nil {
		logger.Error("Failed to load PostgreSQL config", "error", err)
		os.Exit(1)
	}

	logger.Info("Connecting to PostgreSQL database")
	client, err := postgre.NewPostgreSQLClient(postgreConfig, logger)
	if err != nil {
		logger.Error("Postgresql connection error", "error", err)
		os.Exit(1)
	}
	logger.Info("Successful PostgreSQL connection")

	// Redis
	logger.Info("Loading redis config")
	redisConfig := config.NewRedisCondfig()

	logger.Info("Connecting to redis server")
	rdb := redisCache.NewRedisClient(redisConfig)
	logger.Info("Successful redis connection")

	// Kafka Configuration Loading
	logger.Info("Loading Kafka config")
	kafkaCfg := config.NewKafkaConfig()

	// Initialize Kafka Producer
	logger.Info("Connecting to Kafka server", "brokers", kafkaCfg.Brokers, "topic", kafkaCfg.Topic)
	kafkaProducer := kafka.NewProducer(kafkaCfg, logger)

	// segmentio/kafka-go doesn't "connect" immediately;
	// it opens connections lazily from first Publish.
	logger.Info("Kafka producer successfully initialized")

	return logger, client, rdb, kafkaProducer
}
