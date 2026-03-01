package logging

import (
	"context"
	"fmt"
	"time"

	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/core/ports"
	gormlogger "gorm.io/gorm/logger"
)

type gormAdapter struct {
	logger ports.Logger
	config gormlogger.Config
}

// NewGormAdapter bridges the ports.Logger to GORM's requirements
func NewGormAdapter(l ports.Logger, cfg gormlogger.Config) gormlogger.Interface {
	return &gormAdapter{
		logger: l,
		config: cfg,
	}
}

func (a *gormAdapter) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	newConfig := a.config
	newConfig.LogLevel = level
	return &gormAdapter{logger: a.logger, config: newConfig}
}

func (a *gormAdapter) Info(ctx context.Context, msg string, args ...interface{}) {
	if a.config.LogLevel >= gormlogger.Info {
		a.logger.Info(ctx, fmt.Sprintf(msg, args...))
	}
}

func (a *gormAdapter) Warn(ctx context.Context, msg string, args ...interface{}) {
	if a.config.LogLevel >= gormlogger.Warn {
		a.logger.Warn(ctx, fmt.Sprintf(msg, args...))
	}
}

func (a *gormAdapter) Error(ctx context.Context, msg string, args ...interface{}) {
	if a.config.LogLevel >= gormlogger.Error {
		a.logger.Error(ctx, fmt.Sprintf(msg, args...))
	}
}

func (a *gormAdapter) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if a.config.LogLevel <= gormlogger.Silent {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := fc()

	switch {
	case err != nil && a.config.LogLevel >= gormlogger.Error:
		a.logger.Error(ctx, "GORM Query Error", "err", err, "elapsed", elapsed, "rows", rows, "sql", sql)
	case elapsed > a.config.SlowThreshold && a.config.SlowThreshold != 0 && a.config.LogLevel >= gormlogger.Warn:
		a.logger.Warn(ctx, "GORM Slow Query", "elapsed", elapsed, "rows", rows, "sql", sql)
	case a.config.LogLevel >= gormlogger.Info:
		a.logger.Info(ctx, "GORM Query", "elapsed", elapsed, "rows", rows, "sql", sql)
	}
}
