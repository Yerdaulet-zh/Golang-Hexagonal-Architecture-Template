package postgre

import (
	"context"

	persistency "gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/adapters/repository/postgre/persistency/notification"
	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/core/domain"
	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/core/ports"
	"gorm.io/gorm"
)

type NotificationRepo struct {
	db     *gorm.DB
	logger ports.Logger
}

func NewNotificationRepository(db *gorm.DB, logger ports.Logger) *NotificationRepo {
	return &NotificationRepo{
		db:     db,
		logger: logger,
	}
}

func (r *NotificationRepo) Email(ctx context.Context, email string, message string) error {
	if err := gorm.G[persistency.EmailNotification](r.db).Create(ctx, &persistency.EmailNotification{
		Email:   email,
		Message: message,
	}); err != nil {
		r.logger.Error(domain.LogLevelRepository, "Error while creating a record", err)
		return domain.ErrDatabaseInternalError
	}
	return nil
}
