package postgre

import (
	"context"

	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/core/ports"
	"gorm.io/gorm"
)

type NotificationRepo struct {
	db     *gorm.DB
	logger *ports.Logger
}

func NewNotificationRepository(db *gorm.DB, logger *ports.Logger) *NotificationRepo {
	return &NotificationRepo{
		db:     db,
		logger: logger,
	}
}

func (r *NotificationRepo) Notify(ctx context.Context, email string) error {
	return nil
}
