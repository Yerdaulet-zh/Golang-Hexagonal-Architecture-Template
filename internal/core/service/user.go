// Package service provides the core business logic for the application.
// It implements the application's use cases by orchestrating domain
// models and communicating through ports.
package service

import "gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/core/ports"

type NotificationService struct {
	repo ports.Notification
}

func NewUserService(repo ports.Notification) *NotificationService {
	return &NotificationService{repo: repo}
}
