package result

import "time"

type BookingResult struct {
	BookingID int       `json:"booking_id"`
	StartRent time.Time `json:"start_rent"`
	EndRent   time.Time `json:"end_rent"`
	TotalCost int64     `json:"total_cost"`
	Finished  bool      `json:"finished"`
}

type BookingListResult struct {
	BookingID    int64     `json:"booking_id"`
	CustomerName string    `json:"customer_name"`
	CarName      string    `json:"car_name"`
	StartRent    time.Time `json:"start_rent"`
	EndRent      time.Time `json:"end_rent"`
	TotalCost    int64     `json:"total_cost"`
	Finished     bool      `json:"finished"`
}

type BookingDetailsResult struct {
	BookingResult
	Customer CustomerResult `json:"customer"`
	Car      CarResult      `json:"car"`
}

type BookingCarResult struct {
	BookingResult
	Car CarResult `json:"car"`
}
