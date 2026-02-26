// Package domain contains the core business entities and logic.
// This package is independent of any external frameworks or infrastructure.
package domain

import (
	"time"

	"github.com/google/uuid"
)

type EmailNotification struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Email     string
	Message   string
	ID        uuid.UUID
}
