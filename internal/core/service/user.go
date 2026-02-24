// Package service provides the core business logic for the application.
// It implements the application's use cases by orchestrating domain
// models and communicating through ports.
package service

import "gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/core/ports"

type UserService struct {
	repo ports.User
}

func NewUserService(repo ports.User) *UserService {
	return &UserService{repo: repo}
}
