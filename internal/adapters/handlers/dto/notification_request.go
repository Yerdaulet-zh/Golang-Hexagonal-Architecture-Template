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
