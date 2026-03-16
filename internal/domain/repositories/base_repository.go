package repositories

import (
	"context"
)

type BaseRepository[T any] interface {
	FindAll(ctx context.Context, page, limit *int) ([]T, int64, error)
	FindByID(ctx context.Context, id uint) (*T, error)

	Create(ctx context.Context, newCustomer *T) error
	Update(ctx context.Context, id uint, updatedCustomer T) error
	Delete(ctx context.Context, id uint) error
}