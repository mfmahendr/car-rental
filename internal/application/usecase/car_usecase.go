package usecase

import (
	"context"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/mfmahendr/car-rental/internal/application/input"
	"github.com/mfmahendr/car-rental/internal/application/result"
	"github.com/mfmahendr/car-rental/internal/domain"
	"github.com/mfmahendr/car-rental/internal/domain/entities"
	"github.com/mfmahendr/car-rental/internal/domain/repositories"
)

type CarUsecase struct {
	carRepo  repositories.Car
	validate *validator.Validate
}

func NewCarUsecase(v *validator.Validate, r repositories.Car) *CarUsecase {
	return &CarUsecase{validate: v, carRepo: r}
}

func (u *CarUsecase) ListCars(ctx context.Context, in input.PaginationInput) (*result.PageResult[result.CarResult], error) {
	if err := u.validate.Struct(in); err != nil {
		return nil, err
	}

	cars, total, err := u.carRepo.FindAll(ctx, &in.Page, &in.Size)
	totalPages := (int(total) + in.Size - 1) / in.Size
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return &result.PageResult[result.CarResult]{
				Items:      []result.CarResult{},
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

	carResults := make([]result.CarResult, len(cars))
	for i, car := range cars {
		carResults[i] = result.CarResult{
			CarID:     car.CarID,
			Name:      car.Name,
			Stock:     car.Stock,
			DailyRent: car.DailyRent,
		}
	}

	return &result.PageResult[result.CarResult]{
		Items:      carResults,
		TotalItems: int(total),
		TotalPages: totalPages,
		Page:       in.Page,
		Size:       in.Size,
		HasNext:    in.Page < totalPages,
		HasPrev:    in.Page > 1,
	}, nil
}

func (u *CarUsecase) GetByID(ctx context.Context, id uint) (*result.CarResult, error) {
	car, err := u.carRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &result.CarResult{
		CarID:     car.CarID,
		Name:      car.Name,
		Stock:     car.Stock,
		DailyRent: car.DailyRent,
	}, nil
}

func (u *CarUsecase) CreateCar(ctx context.Context, in input.CreateCarInput) (*result.CarResult, error) {
	if err := u.validate.Struct(in); err != nil {
		return nil, err
	}

	newCustomer := &entities.Car{
		Name:      in.Name,
		Stock:     in.Stock,
		DailyRent: in.DailyRent,
	}

	if err := u.carRepo.Create(ctx, newCustomer); err != nil {
		return nil, err
	}

	return &result.CarResult{
		CarID:     newCustomer.CarID,
		Name:      newCustomer.Name,
		Stock:     newCustomer.Stock,
		DailyRent: newCustomer.DailyRent,
	}, nil
}

func (u *CarUsecase) UpdateCar(ctx context.Context, id uint, in input.UpdateCarInput) error {
	if err := u.validate.Struct(in); err != nil {
		return err
	}

	updatedCar := entities.Car{
		CarID:     int(id),
		Name:      in.Name,
		Stock:     in.Stock,
		DailyRent: in.DailyRent,
	}

	return u.carRepo.Update(ctx, id, updatedCar)
}

func (u *CarUsecase) UpdateStock(ctx context.Context, id int64, stock int) error {
	if err := u.validate.Var(stock, "required,min=0"); err != nil {
		return err
	}

	return u.carRepo.UpdateStock(ctx, id, stock)
}

func (u *CarUsecase) DeleteCar(ctx context.Context, id uint) error {
	return u.carRepo.Delete(ctx, id)
}
