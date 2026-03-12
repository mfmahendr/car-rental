package repositories

import (
	"context"
	"time"

	"github.com/mfmahendr/car-rental/internal/domain/entities"
)

type Booking interface {
	BaseRepository[entities.Booking]
	GetBookingsByUserID(ctx context.Context, userID uint) ([]entities.Booking, int64, error)
	GetBookingsByCarID(ctx context.Context, carID uint) ([]entities.Booking, int64, error)
	UpdateRentDate(ctx context.Context, id uint, startRent, endRent time.Time) error
}