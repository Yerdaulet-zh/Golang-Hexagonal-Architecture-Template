package domain

import (
	"time"

	"github.com/google/uuid"
)

type EmailNotification struct {
	ID        uuid.UUID
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
