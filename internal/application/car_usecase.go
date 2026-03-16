package application

import (
	"context"

	"github.com/mfmahendr/car-rental/internal/application/input"
	"github.com/mfmahendr/car-rental/internal/application/result"
)

type CarUsecase interface {
	ListCars(ctx context.Context, in input.PaginationInput) (*result.PageResult[result.CarResult], error)
	GetByID(ctx context.Context, id uint) (*result.CarResult, error)

	CreateCar(ctx context.Context, in input.CreateCarInput) (*result.CarResult, error)
	UpdateCar(ctx context.Context, id uint, in input.UpdateCarInput) error
	UpdateStock(ctx context.Context, id int64, stock int) error
	DeleteCar(ctx context.Context, id uint) error
}
