package service

import (
	"context"
	"service-customer/internal/model"
)

//go:generate go run github.com/vektra/mockery/v2@latest --name=CustomerRepository --output=../../mocks --outpkg=mocks --with-expecter=true
type CustomerRepository interface {
	GetCustomerByID(ctx context.Context, id int64) (*model.Customer, error)
}

type CustomerService struct {
	repo CustomerRepository
}

func NewCustomerService(r CustomerRepository) *CustomerService {
	return &CustomerService{repo: r}
}

func (s *CustomerService) GetCustomerByID(ctx context.Context, id int64) (*model.Customer, error) {
	return s.repo.GetCustomerByID(ctx, id)
}
