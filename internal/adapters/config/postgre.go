package config

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
	"gorm.io/gorm/logger"
)

type DBConfig struct {
	dsn             string
	maxIdleConns    int
	maxOpenConns    int
	connMaxLifetime time.Duration
	connMaxIdleTime time.Duration
	GormLoggerConfig
}

type GormLoggerConfig struct {
	slowThreshold             time.Duration
	logLevel                  logger.LogLevel
	ignoreRecordNotFoundError bool
	colorful                  bool
}

func NewDefaultDBConfig() (*DBConfig, error) {
	return &DBConfig{
		dsn:             os.Getenv("POSTGRES_DSN"), // "host=postgres user=admin password=password dbname=myapp port=5432 sslmode=disable",
		maxIdleConns:    viper.GetInt("database.postgres.MaxIdleConns"),
		maxOpenConns:    viper.GetInt("database.postgres.MaxOpenConns"),
		connMaxLifetime: viper.GetDuration("database.postgres.ConnMaxLifetime"),
		connMaxIdleTime: viper.GetDuration("database.postgres.ConnMaxIdleTime"),
	}, nil
}

func NewGormLoggerConfig() (*GormLoggerConfig, error) {
	var level logger.LogLevel
	switch strings.ToLower(viper.GetString("database.gormLogger.LogLevel")) {
	case "silent":
		level = logger.Silent
	case "error":
		level = logger.Error
	case "warn":
		level = logger.Warn
	case "info":
		level = logger.Info
	default:
		return nil, fmt.Errorf("invalid GORM log level: %s", viper.GetString("database.gormLogger.LogLevel"))
	}

	return &GormLoggerConfig{
		slowThreshold:             viper.GetDuration("database.gormLogger.SlowThreshold"),
		logLevel:                  level,
		ignoreRecordNotFoundError: viper.GetBool("database.gormLogger.IgnoreRecordNotFoundError"),
		colorful:                  viper.GetBool("database.gormLogger.Colorful"),
	}, nil
}

// Getter methods for DBConfig fields
func (c *DBConfig) DSN() string                    { return c.dsn }
func (c *DBConfig) MaxIdleConns() int              { return c.maxIdleConns }
func (c *DBConfig) MaxOpenConns() int              { return c.maxOpenConns }
func (c *DBConfig) ConnMaxLifetime() time.Duration { return c.connMaxLifetime }
func (c *DBConfig) ConnMaxIdleTime() time.Duration { return c.connMaxIdleTime }

// Getter methods for GormLoggerConfig fields
func (l *GormLoggerConfig) SlowThreshold() time.Duration    { return l.slowThreshold }
func (l *GormLoggerConfig) LogLevel() logger.LogLevel       { return l.logLevel }
func (l *GormLoggerConfig) IgnoreRecordNotFoundError() bool { return l.ignoreRecordNotFoundError }
func (l *GormLoggerConfig) Colorful() bool                  { return l.colorful }
