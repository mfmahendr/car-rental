package usecase

import (
	"context"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/mfmahendr/car-rental/internal/application"
	"github.com/mfmahendr/car-rental/internal/application/input"
	"github.com/mfmahendr/car-rental/internal/application/result"
	"github.com/mfmahendr/car-rental/internal/domain"
	"github.com/mfmahendr/car-rental/internal/domain/entities"
	"github.com/mfmahendr/car-rental/internal/domain/repositories"
)

type BookingUsecase struct {
	bookingRepo  repositories.Booking
	carRepo      repositories.Car
	customerRepo repositories.Customer
	validate     *validator.Validate
	transactor   application.Transactor
}

func NewBookingUsecase(v *validator.Validate, t application.Transactor, b repositories.Booking, cr repositories.Car, cs repositories.Customer) *BookingUsecase {
	return &BookingUsecase{bookingRepo: b, carRepo: cr, customerRepo: cs, validate: v, transactor: t}
}

func (u *BookingUsecase) ListBookings(ctx context.Context, in input.PaginationInput) (*result.PageResult[result.BookingListResult], error) {
	if err := u.validate.Struct(in); err != nil {
		return nil, err
	}

	bookings, total, err := u.bookingRepo.FindAll(ctx, &in.Page, &in.Size)
	totalPages := (int(total) + in.Size - 1) / in.Size
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return &result.PageResult[result.BookingListResult]{
				Items:      []result.BookingListResult{},
				TotalItems: int(total),
				TotalPages: totalPages,
				Page:       in.Page,
				Size:       in.Size,
				HasNext:    false,
				HasPrev:    false,
			}, nil
		}
		return nil, err
	}

	// get customer names and car names
	customerNames := make(map[uint]string)
	customerIDs := make([]uint, 0, len(bookings))
	carNames := make(map[uint]string)
	carIDs := make([]uint, 0, len(bookings))
	for _, b := range bookings {
		customerNames[uint(b.CustomerID)] = ""
		customerIDs = append(customerIDs, uint(b.CustomerID))
		carNames[uint(b.CarID)] = ""
		carIDs = append(carIDs, uint(b.CarID))
	}

	if len(customerNames) > 0 {
		customers, err := u.customerRepo.FindByIDs(ctx, customerIDs)
		if err != nil {
			return nil, err
		}

		for _, c := range customers {
			customerNames[uint(c.CustomerID)] = c.Name
		}
	}

	if len(carNames) > 0 {
		cars, err := u.carRepo.FindByIDs(ctx, carIDs)
		if err != nil {
			return nil, err
		}

		for _, c := range cars {
			carNames[uint(c.CarID)] = c.Name
		}
	}

	bookingResuts := make([]result.BookingListResult, len(bookings))
	for i, b := range bookings {
		bookingResuts[i] = result.BookingListResult{
			BookingID:    int64(b.BookingID),
			CustomerName: customerNames[uint(b.CustomerID)],
			CarName:      carNames[uint(b.CarID)],
			StartRent:    b.StartRent,
			EndRent:      b.EndRent,
			TotalCost:    b.TotalCost,
			Finished:     b.Finished,
		}
	}

	return &result.PageResult[result.BookingListResult]{
		Items:      bookingResuts,
		TotalItems: int(total),
		TotalPages: totalPages,
		Page:       in.Page,
		Size:       in.Size,
		HasNext:    in.Page < totalPages,
		HasPrev:    in.Page > 1,
	}, nil
}

func (u *BookingUsecase) GetCustomerBookingHistory(ctx context.Context, customerID uint) ([]result.BookingCarResult, error) {
	bookings, _, err := u.bookingRepo.GetBookingsByUserID(ctx, customerID)
	if err != nil {
		return nil, err
	}

	carIDsMap := make(map[uint]struct{})
	for _, b := range bookings {
		carIDsMap[uint(b.CarID)] = struct{}{}
	}

	var ids []uint
	for id := range carIDsMap {
		ids = append(ids, id)
	}

	cars, err := u.carRepo.FindByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}

	carMap := make(map[uint]entities.Car)
	for _, c := range cars {
		carMap[uint(c.CarID)] = c
	}

	res := make([]result.BookingCarResult, len(bookings))
	for i, b := range bookings {
		res[i] = result.BookingCarResult{
			BookingResult: result.BookingResult{
				BookingID: b.BookingID,
				StartRent: b.StartRent,
				EndRent:   b.EndRent,
				TotalCost: b.TotalCost,
				Finished:  b.Finished,
			},
			Car: result.CarResult{
				CarID:     carMap[uint(b.CarID)].CarID,
				Name:      carMap[uint(b.CarID)].Name,
				Stock:     carMap[uint(b.CarID)].Stock,
				DailyRent: carMap[uint(b.CarID)].DailyRent,
			},
		}
	}

	return res, nil
}

func (u *BookingUsecase) FindByID(ctx context.Context, id uint) (*result.BookingResult, error) {
	booking, err := u.bookingRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &result.BookingResult{
		BookingID: booking.BookingID,
		StartRent: booking.StartRent,
		EndRent:   booking.EndRent,
		TotalCost: booking.TotalCost,
		Finished:  booking.Finished,
	}, nil
}

func (u *BookingUsecase) CreateBooking(ctx context.Context, in input.CreateBookingInput) (*result.BookingResult, error) {
	if err := u.validate.Struct(in); err != nil {
		return nil, err
	}

	var newBooking *entities.Booking

	err := u.transactor.WithTransaction(ctx, func(txCtx context.Context) error {
		// Check car and stock
		car, err := u.carRepo.FindByID(txCtx, uint(in.CarID))
		if err != nil {
			return err
		}
		if car.Stock < 1 {
			return application.ErrCarOutOfStock
		}

		// save the booking data
		newBooking = &entities.Booking{
			CustomerID: in.CustomerID,
			CarID:      in.CarID,
			StartRent:  in.StartRent,
			EndRent:    in.EndRent,
			TotalCost:  int64(in.EndRent.Sub(in.StartRent).Hours()/24) * car.DailyRent,
			Finished:   false,
		}

		if err := u.bookingRepo.Create(txCtx, newBooking); err != nil {
			return err
		}

		// reduce the stock of the car
		if err := u.carRepo.UpdateStock(txCtx, int64(car.CarID), -1); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &result.BookingResult{
		BookingID: newBooking.BookingID,
		StartRent: newBooking.StartRent,
		EndRent:   newBooking.EndRent,
		TotalCost: newBooking.TotalCost,
		Finished:  newBooking.Finished,
	}, nil
}

func (u *BookingUsecase) FinishBooking(ctx context.Context, id uint) error {
	return u.transactor.WithTransaction(ctx, func(txCtx context.Context) error {
		booking, err := u.bookingRepo.FindByID(txCtx, id)
		if err != nil {
			return err
		}

		if booking.Finished {
			return application.ErrBookingFinished
		}

		// update booking status
		booking.Finished = true
		if err := u.bookingRepo.Update(txCtx, id, *booking); err != nil {
			return err
		}

		// add stock back to the car inventory
		if err := u.carRepo.UpdateStock(txCtx, int64(booking.CarID), 1); err != nil {
			return err
		}

		return nil
	})
}

func (u *BookingUsecase) UpdateRentDate(ctx context.Context, id uint, in input.UpdateBookingRentDateInput) error {
	if err := u.validate.Struct(in); err != nil {
		return err
	}

	if !in.EndRent.After(in.StartRent) {
		return application.ErrInvalidRentDate
	}

	booking, err := u.bookingRepo.FindByID(ctx, id)
	if err != nil {
		return application.ErrBookingNotFound
	}

	if booking.Finished {
		return application.ErrBookingFinished
	}

	car, err := u.carRepo.FindByID(ctx, uint(booking.CarID))
	if err != nil {
		return application.ErrCarNotFound
	}

	days := int64(in.EndRent.Sub(in.StartRent).Hours() / 24)
	if days == 0 {
		days = 1
	}
	updatedBooking := *booking
	updatedBooking.StartRent = in.StartRent
	updatedBooking.EndRent = in.EndRent
	updatedBooking.TotalCost = days * car.DailyRent

	return u.bookingRepo.Update(ctx, id, updatedBooking)
}

func (u *BookingUsecase) DeleteBooking(ctx context.Context, id uint) error {
	return u.bookingRepo.Delete(ctx, id)
}
