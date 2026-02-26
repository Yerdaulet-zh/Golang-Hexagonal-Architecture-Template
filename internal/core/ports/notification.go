// Package ports defines the interfaces that bridge the core logic with external adapters.
package ports

import "context"

type Notification interface {
	Email(ctx context.Context, email string, message string) error
}

type Validator interface {
	ValidateEmailHost(email string) error
}
