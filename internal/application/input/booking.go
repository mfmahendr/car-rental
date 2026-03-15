package input

import "time"

type CreateBookingInput struct {
	CustomerID int       `json:"customer_id" validate:"required"`
	CarID      int       `json:"car_id" validate:"required"`
	StartRent  time.Time `json:"start_rent" validate:"required,ltefield=EndRent"`
	EndRent    time.Time `json:"end_rent" validate:"required,gtefield=StartRent"`
}

type UpdateBookingRentDateInput struct {
	StartRent time.Time `json:"start_rent" validate:"omitempty,ltefield=EndRent"`
	EndRent   time.Time `json:"end_rent" validate:"omitempty,gtefield=StartRent"`
}
