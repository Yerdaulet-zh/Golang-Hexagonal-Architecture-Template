// Package postgre provides the PostgreSQL database adapter implementation.
// It handles connection management, GORM initialization, and implements
// the ports.Database interface for the application core.
package postgre

import (
	"context"
	"fmt"

	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/adapters/config"
	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/adapters/logging" // Added import for your logger adapter
	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/core/domain"
	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/core/ports"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger" // Aliased for clarity
	"gorm.io/gorm/schema"
)

type Client struct {
	DB *gorm.DB
}

// NewPostgreSQLClient initializes the GORM client using the provided configuration and custom logger.
func NewPostgreSQLClient(cfg *config.DBConfig, appLogger ports.Logger) (*Client, error) {
	if cfg.DSN() == "" {
		return nil, domain.ErrInvalidDSN
	}

	// Pass the appLogger down to the opening logic
	db, err := openPostgreSQLDB(cfg, appLogger)
	if err != nil {
		appLogger.Error(context.TODO(), domain.LogLevelRepository, "Error while opening a new PostgreSQL database", "error", err)
		return nil, domain.ErrPostgreSQLOpenDB
	}
	return &Client{DB: db}, nil
}

func openPostgreSQLDB(cfg *config.DBConfig, appLogger ports.Logger) (*gorm.DB, error) {
	dialector := postgres.Open(cfg.DSN())

	// Use the custom logging adapter instead of the standard gorm logger
	// This routes GORM internal logs through logging.NewGormAdapter
	gormLogger := logging.NewGormAdapter(
		appLogger,
		gormlogger.Config{
			SlowThreshold:             cfg.SlowThreshold(),
			LogLevel:                  cfg.LogLevel(),
			IgnoreRecordNotFoundError: cfg.IgnoreRecordNotFoundError(),
			Colorful:                  cfg.Colorful(),
		},
	)

	// GORM configuration
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
		Logger:         gormLogger, // Plug in the custom adapter here
	}

	// Open database connection
	db, err := gorm.Open(dialector, gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to open PostgreSQL connection: %w", err)
	}

	// Enable tracing (otelgorm will use the context from requests to link spans)
	if err = db.Use(otelgorm.NewPlugin()); err != nil {
		return nil, fmt.Errorf("failed to register otelgorm plugin: %w", err)
	}

	// Configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get generic database object: %w", err)
	}

	// Ensure database connection is reachable
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("database is unreachable: %w", err)
	}

	// Connection pool settings
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns())
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns())
	sqlDB.SetConnMaxLifetime(cfg.ConnMaxLifetime())
	sqlDB.SetConnMaxIdleTime(cfg.ConnMaxIdleTime())

	return db, nil
}

// Ping checks the health of the database connection.
func (c *Client) Ping(ctx context.Context) error {
	sqlDB, err := c.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.PingContext(ctx)
}

// Close terminates the database connection pool.
func (c *Client) Close() error {
	sqlDB, err := c.DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB for closing: %w", err)
	}
	return sqlDB.Close()
}

// GetGormDB returns the underlying GORM database instance.
func (c *Client) GetGormDB() *gorm.DB {
	return c.DB
}
