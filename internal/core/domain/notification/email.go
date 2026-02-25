// Package domain contains the core business entities and logic.
// This package is independent of any external frameworks or infrastructure.
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
