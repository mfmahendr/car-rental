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

type CustomerUsecase struct {
	customerRepo repositories.Customer
	validate     *validator.Validate
}

func NewCustomerUsecase(v *validator.Validate, r repositories.Customer) *CustomerUsecase {
	return &CustomerUsecase{validate: v, customerRepo: r}
}

func (u *CustomerUsecase) ListCustomers(ctx context.Context, in input.PaginationInput) (*result.PageResult[result.CustomerResult], error) {
	if err := u.validate.Struct(in); err != nil {
		return nil, err
	}

	customers, total, err := u.customerRepo.FindAll(ctx, &in.Page, &in.Size)
	totalPages := (int(total) + in.Size - 1) / in.Size
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return &result.PageResult[result.CustomerResult]{
				Items:      []result.CustomerResult{},
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

	customerResults := make([]result.CustomerResult, len(customers))
	for i, customer := range customers {
		customerResults[i] = result.CustomerResult{
			CustomerID:  customer.CustomerID,
			Name:        customer.Name,
			NIK:         customer.NIK,
			PhoneNumber: customer.PhoneNumber,
		}
	}

	return &result.PageResult[result.CustomerResult]{
		Items:      customerResults,
		TotalItems: int(total),
		TotalPages: totalPages,
		Page:       in.Page,
		Size:       in.Size,
		HasNext:    in.Page < totalPages,
		HasPrev:    in.Page > 1,
	}, nil
}

func (u *CustomerUsecase) GetByID(ctx context.Context, id uint) (*result.CustomerResult, error) {
	customer, err := u.customerRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &result.CustomerResult{
		CustomerID:  customer.CustomerID,
		Name:        customer.Name,
		NIK:         customer.NIK,
		PhoneNumber: customer.PhoneNumber,
	}, nil
}

func (u *CustomerUsecase) CreateCustomer(ctx context.Context, in input.CreateCustomerInput) (*result.CustomerResult, error) {
	if err := u.validate.Struct(in); err != nil {
		return nil, err
	}

	newCustomer := &entities.Customer{
		Name:        in.Name,
		NIK:         in.NIK,
		PhoneNumber: in.PhoneNumber,
	}

	if err := u.customerRepo.Create(ctx, newCustomer); err != nil {
		return nil, err
	}

	return &result.CustomerResult{
		CustomerID:  newCustomer.CustomerID,
		Name:        newCustomer.Name,
		NIK:         newCustomer.NIK,
		PhoneNumber: newCustomer.PhoneNumber,
	}, nil
}

func (u *CustomerUsecase) UpdateCustomer(ctx context.Context, id uint, in input.UpdateCustomerInput) error {
	if err := u.validate.Struct(in); err != nil {
		return err
	}

	updatedCustomer := entities.Customer{
		CustomerID:  int(id),
		Name:        in.Name,
		NIK:         in.NIK,
		PhoneNumber: in.PhoneNumber,
	}

	return u.customerRepo.Update(ctx, id, updatedCustomer)
}

func (u *CustomerUsecase) DeleteCustomer(ctx context.Context, id uint) error {
	return u.customerRepo.Delete(ctx, id)
}
