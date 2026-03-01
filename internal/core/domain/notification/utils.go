package domain

import (
	"context"
	"fmt"

	"github.com/badoux/checkmail"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)

type Validator struct{}

func NewNotificationValidator() *Validator {
	return &Validator{}
}

func (v *Validator) ValidateEmailHost(ctx context.Context, email string) error {
	_, span := otel.Tracer("domain").Start(ctx, "Validator.ValidateEmailHost")
	defer span.End()

	if err := checkmail.ValidateHost(email); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "DNS host validation failed")
		return fmt.Errorf("email service validation error: %s", err.Error())
	}
	return nil
}
