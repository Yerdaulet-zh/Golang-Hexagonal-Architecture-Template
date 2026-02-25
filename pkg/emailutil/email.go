// Package pkg provides general-purpose utilities and helper functions
// that are shared across the application but are not tied to any
// specific business domain or infrastructure.
package pkg

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func IsValidEmailFormat(email string) error {
	if err := checkmail.ValidateFormat(email); err != nil {
		return fmt.Errorf("Email syntax validation error: %s", err.Error())
	}
	return nil
}
