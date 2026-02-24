package service

import "gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/core/ports"

type UserService struct {
	userRepo *ports.User
}

func NewUserService(repo *ports.User) *UserService {
	return &UserService{
		userRepo: repo,
	}
}
