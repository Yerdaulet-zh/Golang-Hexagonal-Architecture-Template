// Package ports defines the interfaces that bridge the core logic with external adapters.
package ports

type User interface {
	// Create(ctx context.Context, u *domain.User) error
	// GetByID(ctx context.Context, id string) (*domain.User)
}
