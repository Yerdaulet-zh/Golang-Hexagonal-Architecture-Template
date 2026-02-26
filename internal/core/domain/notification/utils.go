package domain

import (
	"fmt"

	"github.com/badoux/checkmail"
)

type Validator struct{}

func NewNotificationValidator() *Validator {
	return &Validator{}
}

func (v *Validator) ValidateEmailHost(email string) error {
	if err := checkmail.ValidateHost(email); err != nil {
		return fmt.Errorf("Email service validation error: %s", err.Error())
	}
	return nil
}
