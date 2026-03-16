package application

import (
	"context"

	"github.com/mfmahendr/car-rental/internal/application/input"
	"github.com/mfmahendr/car-rental/internal/application/result"
)

type CustomerUsecase interface {
	ListCustomers(ctx context.Context, in input.PaginationInput) (*result.PageResult[result.CustomerResult], error)
	GetByID(ctx context.Context, id uint) (*result.CustomerResult, error)

	CreateCustomer(ctx context.Context, in input.CreateCustomerInput) (*result.CustomerResult, error)
	UpdateCustomer(ctx context.Context, id uint, in input.UpdateCustomerInput) error
	DeleteCustomer(ctx context.Context, id uint) error
}
