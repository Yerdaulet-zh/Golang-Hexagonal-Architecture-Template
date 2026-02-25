package persistency

import (
	"time"

	"github.com/google/uuid"
	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/core/domain"
	"gorm.io/gorm"
)

type EmailNotification struct {
	ID        uuid.UUID      `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Email     string         `gorm:"type:varchar(100);not null"`
	CreatedAt time.Time      `gorm:"type:timestamptz;default:now();not null"`
	UpdatedAt time.Time      `gorm:"type:timestamptz;default:now();not null"`
	DeletedAt gorm.DeletedAt `gorm:"type:timestamptz;index"`
}

func (n *EmailNotification) ToDomain() *domain.EmailNotification {
	return &domain.EmailNotification{
		ID:        n.ID,
		Email:     n.Email,
		CreatedAt: n.CreatedAt,
		UpdatedAt: n.UpdatedAt,
		DeletedAt: n.DeletedAt,
	}
}
