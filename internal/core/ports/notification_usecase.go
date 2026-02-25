package ports

import "context"

type NotificationUseCase interface {
	Email(ctx context.Context, email string) error
}
