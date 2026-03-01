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
	"time"

	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/cmd/servers"
	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/adapters/broker/kafka"
	redisCache "gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/adapters/cache/redis"
	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/adapters/config"
	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/adapters/logging"
	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/adapters/repository/postgre"
	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/adapters/tracing"

	domain "gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/core/domain/notification"
	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/core/ports"
	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/core/service"
	"go.opentelemetry.io/otel/sdk/trace"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	logger, tracer, client, rdb, _ := loadComponents(ctx)

	if err := run(ctx, logger, tracer, client, rdb); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, logger ports.Logger, tracer *trace.TracerProvider, client ports.Database, rdb ports.Redis) error {
	defer func() {
		logger.Info(ctx, "Closing infrastructure connections...")
		if err := client.Close(); err != nil {
			logger.Error(ctx, "Postgre close error", "error", err)
		}
		if err := rdb.Close(); err != nil {
			logger.Error(ctx, "Redis close error", "error", err)
		}
		// Gracefull shutdown of tracer
		shutdownCtx, cancelTracer := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelTracer()
		if err := tracer.Shutdown(shutdownCtx); err != nil {
			logger.Error(ctx, "Failed to shutdown tracer", "error", err)
		}
		logger.Info(ctx, "Done")
	}()

	logger.Info(ctx, "Loading HTTP Server config")
	httpConfig := config.NewHttpConfig()
	logger.Info(ctx, "Successfully loaded HTTP Server config")

	// Notification Domain
	notificationValidator := domain.NewNotificationValidator()
	notificationRepo := postgre.NewNotificationRepository(client.GetGormDB(), logger)
	notificationService := service.NewNotificationService(notificationRepo, logger, notificationValidator)

	mapBusinessHandler := servers.MapBusinessRoutes(logger, tracer, rdb, notificationService)
	mapManagementRoutes := servers.MapManagementRoutes(logger, client)

	go func() {
		if err := servers.Run(ctx, logger, mapManagementRoutes, httpConfig.HttpManagementAddr, httpConfig.GracefullShutdown, "Management"); err != nil {
			logger.Error(ctx, "HTTP Management server error while shutting down", "error", err)
		}
	}()

	if err := servers.Run(ctx, logger, mapBusinessHandler, httpConfig.HttpBusinessAddr, httpConfig.GracefullShutdown, "Business"); err != nil {
		logger.Error(ctx, "HTTP Business server error while shutting down", "error", err)
		os.Exit(1)
	}
	logger.Info(ctx, "Application exited cleanly")
	return nil
}

func loadComponents(ctx context.Context) (ports.Logger, *trace.TracerProvider, *postgre.Client, ports.Redis, ports.Broker) {
	// Configuration
	cfg, err := config.NewLoggingConfig()
	if err != nil {
		log.Fatalf("Error initializing config: %v", err)
	}

	// Logger
	logger := logging.NewLogger(cfg)
	logger.Info(ctx, "Logging successfully configured to use the adapter: ", cfg.Adapter())

	// Tracer
	tracer, err := tracing.InitTracer()
	if err != nil {
		logger.Error(ctx, "Failed to init OTel tracer", err)
		os.Exit(1)
	}

	// PostgreSQL
	logger.Info(ctx, "Loading PostgreSQL config")
	postgreConfig, err := config.NewDefaultDBConfig()
	if err != nil {
		logger.Error(ctx, "Failed to load PostgreSQL config", "error", err)
		os.Exit(1)
	}

	logger.Info(ctx, "Connecting to PostgreSQL database")
	client, err := postgre.NewPostgreSQLClient(postgreConfig, logger)
	if err != nil {
		logger.Error(ctx, "Postgresql connection error", "error", err)
		os.Exit(1)
	}
	logger.Info(ctx, "Successful PostgreSQL connection")

	// Redis
	logger.Info(ctx, "Loading redis config")
	redisConfig := config.NewRedisCondfig()

	logger.Info(ctx, "Connecting to redis server")
	rdb := redisCache.NewRedisClient(logger, redisConfig)
	logger.Info(ctx, "Successful redis connection")

	// Kafka Configuration Loading
	logger.Info(ctx, "Loading Kafka config")
	kafkaCfg := config.NewKafkaConfig()

	// Initialize Kafka Producer
	logger.Info(ctx, "Connecting to Kafka server", "brokers", kafkaCfg.Brokers, "topic", kafkaCfg.Topic)
	kafkaProducer := kafka.NewProducer(kafkaCfg, logger)

	// segmentio/kafka-go doesn't "connect" immediately;
	// it opens connections lazily from first Publish.
	logger.Info(ctx, "Kafka producer successfully initialized")
	return logger, tracer, client, rdb, kafkaProducer
}
