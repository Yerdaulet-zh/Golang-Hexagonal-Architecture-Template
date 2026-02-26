// Package dto (Data Transfer Objects) defines the schema for data entering
// the system via external primary adapters, such as HTTP handlers.
// It includes logic for request-scoped sanitization and validation tags.
package dto

import "strings"

type EmailNotification struct {
	Email   string `json:"email" validate:"required,email"`
	Message string `json:"message" validate:"required,max=255"`
}

func (r *EmailNotification) Sanitize() {
	r.Email = strings.TrimSpace(r.Email)
	r.Message = strings.TrimSpace(r.Message)
}
