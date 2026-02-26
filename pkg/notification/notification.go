// Package pkg provides general-purpose utilities and helper functions
// that are shared across the application but are not tied to any
// specific business domain or infrastructure.
package notificationutil

import (
	"fmt"
	"unicode/utf8"

	"github.com/badoux/checkmail"
)

func IsValidEmailFormat(email string) error {
	if err := checkmail.ValidateFormat(email); err != nil {
		return fmt.Errorf("email syntax validation error: %s", err.Error())
	}
	return nil
}

func IsValidLenght(message string) bool {
	return utf8.RuneCountInString(message) <= 255
}
