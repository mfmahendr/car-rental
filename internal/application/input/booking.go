package input

import "time"

type CreateBookingInput struct {
	CustomerID int       `validate:"required"`
	CarID      int       `validate:"required"`
	StartRent  time.Time `validate:"required,ltefield=EndRent"`
	EndRent    time.Time `validate:"required,gtefield=StartRent"`
}

type UpdateBookingRentDateInput struct {
	StartRent time.Time `validate:"omitempty,ltefield=EndRent"`
	EndRent   time.Time `validate:"omitempty,gtefield=StartRent"`
}
