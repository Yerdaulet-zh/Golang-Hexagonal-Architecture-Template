// Package service provides the core business logic for the application.
// It implements the application's use cases by orchestrating domain
// models and communicating through ports.
package service

import (
	"context"

	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/core/domain"
	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/core/ports"
	pkg "gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/pkg/emailutil"
)

type NotificationService struct {
	repo      ports.Notification
	logger    ports.Logger
	validator ports.Validator
}

func NewUserService(repo ports.Notification, logger ports.Logger, validator ports.Validator) *NotificationService {
	return &NotificationService{
		repo:      repo,
		logger:    logger,
		validator: validator,
	}
}

func (n *NotificationService) Email(ctx context.Context, email string) error {
	if err := pkg.IsValidEmailFormat(email); err != nil {
		n.logger.Error(domain.LogLevelService, "Error while validation Email Format:", err)
		return domain.ErrInvalidEmailFormat
	}
	if err := n.validator.ValidateEmailHost(email); err != nil {
		n.logger.Error(domain.LogLevelService, "Error while validation Email Host:", err)
		return domain.ErrInvalidEmailHost
	}
	return n.repo.Email(ctx, email)
}
