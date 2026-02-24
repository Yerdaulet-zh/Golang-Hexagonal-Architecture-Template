package postgre

import (
	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/core/ports"
	"gorm.io/gorm"
)

type UserRepo struct {
	db     *gorm.DB
	logger *ports.Logger
}

func NewUserRepository(db *gorm.DB, logger *ports.Logger) *UserRepo {
	return &UserRepo{
		db:     db,
		logger: logger,
	}
}
