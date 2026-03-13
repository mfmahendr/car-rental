package application

import "errors"

var (
	ErrCarNotFound     = errors.New("car not found")
	ErrCarOutOfStock   = errors.New("car out of stock")
	ErrBookingNotFound = errors.New("booking not found")
	ErrBookingFinished = errors.New("booking already finished")
	ErrInvalidRentDate = errors.New("invalid rent date")
)
