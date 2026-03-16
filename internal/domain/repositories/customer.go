package repositories

import (
	"context"

	"github.com/mfmahendr/car-rental/internal/domain/entities"
)

type Customer interface {
	BaseRepository[entities.Customer]
	FindByIDs(ctx context.Context, ids []uint) ([]entities.Customer, error)
}
