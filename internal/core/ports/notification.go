// Package ports defines the interfaces that bridge the core logic with external adapters.
package ports

import "context"

type Notification interface {
	Notify(ctx context.Context, email string) error
}
