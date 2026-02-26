package dto

type EmailNotification struct {
	Email   string `json:"email" validate:"required,email"`
	Message string `json:"message" validate:"required,max=255"`
}
