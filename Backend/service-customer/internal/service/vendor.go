package service

import (
	"context"
	"service-customer/internal/model"
)

//go:generate go run github.com/vektra/mockery/v2@latest --name=VendorRepository --output=../../mocks --outpkg=mocks --with-expecter=true
type VendorRepository interface {
	GetVendorByID(ctx context.Context, id int64) (*model.Vendor, error)
}

type VendorService struct {
	repo VendorRepository
}

func NewVendorService(r VendorRepository) *VendorService {
	return &VendorService{repo: r}
}

func (s *VendorService) GetVendorByID(ctx context.Context, id int64) (*model.Vendor, error) {
	return s.repo.GetVendorByID(ctx, id)
}
