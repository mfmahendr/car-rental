package repositories

import (
	"context"

	"github.com/mfmahendr/car-rental/internal/domain/entities"
)

type Car interface {
	BaseRepository[entities.Car]
	FindByIDs(ctx context.Context, ids []uint) ([]entities.Car, error)
	UpdateStock(ctx context.Context, id int64, stockChanges int) error
}
