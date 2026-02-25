// Package persistency contains GORM model implementations for database persistence.
// It acts as a data mapping layer between the physical database schema and
// the internal domain entities, facilitating the conversion between them.
package persistency

import (
	"time"

	"github.com/google/uuid"
	domain "gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/core/domain/notification"
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
	var deletedAt *time.Time
	if n.DeletedAt.Valid {
		deletedAt = &n.DeletedAt.Time
	}
	return &domain.EmailNotification{
		ID:        n.ID,
		Email:     n.Email,
		CreatedAt: n.CreatedAt,
		UpdatedAt: n.UpdatedAt,
		DeletedAt: deletedAt,
	}
}
