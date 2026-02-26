// Package handlers implements the primary adapters for the application's transport layer.
// It is responsible for intercepting incoming HTTP requests, decoding and validating
// user input (DTOs), and invoking the appropriate core business services.
package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/adapters/handlers/common"
	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/adapters/handlers/dto"
	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/core/ports"
)

type Notification struct {
	service  ports.NotificationUseCase
	logger   ports.Logger
	validate *validator.Validate
}

func NewNotificationHandler(service ports.NotificationUseCase, logger ports.Logger) *Notification {
	return &Notification{
		service:  service,
		logger:   logger,
		validate: validator.New(),
	}
}

func (h *Notification) EmailNotification(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusBadRequest)
		return
	}
	var req dto.EmailNotification
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid Request Payload", http.StatusBadRequest)
		return
	}
	req.Sanitize()
	if err := h.validate.Struct(req); err != nil {
		http.Error(w, "invalid Requets Payload", http.StatusBadRequest)
		return
	}
	ctx := r.Context()
	if err := h.service.Email(ctx, req.Email, req.Message); err != nil {
		common.MapErrorToResponse(w, err)
		return
	}

	common.WriteSuccess(w, http.StatusCreated, "the email has been sent successfully", nil)
}
