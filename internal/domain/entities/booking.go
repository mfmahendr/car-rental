package entities

import (
	"time"
)

type Booking struct {
	BookingID  int
	CustomerID int
	CarID      int
	StartRent  time.Time
	EndRent    time.Time
	TotalCost  int64
	Finished   bool
}
