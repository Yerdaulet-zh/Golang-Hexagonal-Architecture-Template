package domain

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func IsValidHost(email string) error {
	if err := checkmail.ValidateHost(email); err != nil {
		return fmt.Errorf("Email service validation error: %s", err.Error())
	}
	return nil
}
