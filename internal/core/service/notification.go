// Package service provides the core business logic for the application.
// It implements the application's use cases by orchestrating domain
// models and communicating through ports.
package service

import (
	"context"

	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/core/domain"
	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/core/ports"
	notificationutil "gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/pkg/notification"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
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
	ctx, span := otel.Tracer("notification-service").Start(ctx, "NotificationService.Email")
	defer span.End()

	span.SetAttributes(
		attribute.String("notification.email", email),
	)

	if err := notificationutil.IsValidEmailFormat(email); err != nil {
		n.recordError(ctx, span, "Invalid Email Format", err)
		return domain.ErrInvalidEmailFormat
	}

	// if err := n.validator.ValidateEmailHost(ctx, email); err != nil {
	// 	n.recordError(ctx, span, "Invalid Email Host", err)
	// 	return domain.ErrInvalidEmailHost
	// }

	if !notificationutil.IsValidLenght(message) {
		n.recordError(ctx, span, "Invalid Message Length", domain.ErrInvalidMessageLenght)
		return domain.ErrInvalidMessageLenght
	}

	return n.repo.Email(ctx, email, message)
}

func (n *NotificationService) recordError(ctx context.Context, span trace.Span, msg string, err error) {
	span.RecordError(err)
	span.SetStatus(codes.Error, msg)
	n.logger.Error(ctx, domain.LogLevelService, msg, err)
}
