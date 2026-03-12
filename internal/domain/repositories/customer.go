package repositories

import "github.com/mfmahendr/car-rental/internal/domain/entities"

type Customer interface {
	BaseRepository[entities.Customer]
}