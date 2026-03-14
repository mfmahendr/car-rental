package request

import "time"

type CreateBooking struct {
	CustomerID int       `json:"customer_id" form:"customer_id"`
	CarID      int       `json:"car_id" form:"car_id"`
	StartRent  time.Time `json:"start_rent" form:"start_rent"`
	EndRent    time.Time `json:"end_rent" form:"end_rent"`
}

type UpdateBookingRentDate struct {
	StartRent time.Time `json:"start_rent" form:"start_rent"`
	EndRent   time.Time `json:"end_rent" form:"end_rent"`
}

