package repositories

import (
	"github.com/mfmahendr/car-rental/internal/domain/entities"
)

type Car interface {
	BaseRepository[entities.Car]
}