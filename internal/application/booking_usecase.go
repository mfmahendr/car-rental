package application

import (
	"context"

	"github.com/mfmahendr/car-rental/internal/application/input"
	"github.com/mfmahendr/car-rental/internal/application/result"
)

type BookingUsecase interface {
	ListBookings(ctx context.Context, in input.PaginationInput) (*result.PageResult[result.BookingListResult], error)
	FindByID(ctx context.Context, id uint) (*result.BookingResult, error)
	GetCustomerBookingHistory(ctx context.Context, customerID uint) ([]result.BookingCarResult, error)

	CreateBooking(ctx context.Context, in input.CreateBookingInput) (*result.BookingResult, error)
	UpdateRentDate(ctx context.Context, id uint, in input.UpdateBookingRentDateInput) error
	FinishBooking(ctx context.Context, id uint) error
	DeleteBooking(ctx context.Context, id uint) error
}
