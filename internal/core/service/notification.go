// Package service provides the core business logic for the application.
// It implements the application's use cases by orchestrating domain
// models and communicating through ports.
package service

import (
	"context"

	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/core/domain"
	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/core/ports"
	notificationutil "gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/pkg/notification"
)

type NotificationService struct {
	repo      ports.Notification
	logger    ports.Logger
	validator ports.Validator
}

func NewNotificationService(repo ports.Notification, logger ports.Logger, validator ports.Validator) *NotificationService {
	return &NotificationService{
		repo:      repo,
		logger:    logger,
		validator: validator,
	}
}

func (n *NotificationService) Email(ctx context.Context, email string, message string) error {
	if err := notificationutil.IsValidEmailFormat(email); err != nil {
		n.logger.Error(domain.LogLevelService, "Error while validation Email Format:", err)
		return domain.ErrInvalidEmailFormat
	}
	if err := n.validator.ValidateEmailHost(email); err != nil {
		n.logger.Error(domain.LogLevelService, "Error while validation Email Host:", err)
		return domain.ErrInvalidEmailHost
	}
	if !notificationutil.IsValidLenght(message) {
		n.logger.Error(domain.LogLevelService, "Error while validationg Notification Message lenght", domain.ErrInvalidMessageLenght)
		return domain.ErrInvalidMessageLenght
	}

	return n.repo.Email(ctx, email, message)
}
