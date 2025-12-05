package service

import (
	"context"
	"service-customer/internal/model"
)

//go:generate go run github.com/vektra/mockery/v2@latest --name=ProductRepository --output=../../mocks --outpkg=mocks --with-expecter=true
type ProductRepository interface {
	GetProductByID(ctx context.Context, id int64) (*model.Product, error)
	GetProducts(ctx context.Context, offset, limit int, minPrice, maxPrice float64, vendorID int64, queryParam, sortParam string) ([]model.Product, int, error)
}

//go:generate go run github.com/vektra/mockery/v2@latest --name=ProductVendorService --output=../../mocks --outpkg=mocks --with-expecter=true
type ProductVendorService interface {
	GetVendorByID(ctx context.Context, id int64) (*model.Vendor, error)
}

//go:generate go run github.com/vektra/mockery/v2@latest --name=ProductReviewService --output=../../mocks --outpkg=mocks --with-expecter=true
type ProductReviewService interface {
	GetReviewStats(ctx context.Context, productID int64) (int, float64, error)
}

type ProductService struct {
	repo          ProductRepository
	vendorService ProductVendorService
	reviewService ProductReviewService
}

func NewProductService(r ProductRepository, vendorService ProductVendorService, reviewService ProductReviewService) *ProductService {
	return &ProductService{
		repo:          r,
		vendorService: vendorService,
		reviewService: reviewService,
	}
}

func (s *ProductService) GetProductByID(ctx context.Context, productID int64) (*model.Product, error) {
	return s.repo.GetProductByID(ctx, productID)
}

func (s *ProductService) GetFullProductData(ctx context.Context, productID int64) (*model.Product, *model.Vendor, int, float64, error) {
	product, err := s.repo.GetProductByID(ctx, productID)
	if err != nil {
		return nil, nil, 0, 0, err
	}

	vendor, err := s.vendorService.GetVendorByID(ctx, product.VendorID)
	if err != nil {
		return nil, nil, 0, 0, err
	}

	//TODO: добавить обработку ошибки (пока нет логирования с ней нечего делать)
	count, avg, _ := s.reviewService.GetReviewStats(ctx, productID)

	return product, vendor, count, avg, nil
}

func (s *ProductService) GetProducts(ctx context.Context, page, size int, minPrice, maxPrice float64, vendorID int64, query, sort string) ([]model.Product, int, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 20
	}
	offset := (page - 1) * size
	return s.repo.GetProducts(ctx, offset, size, minPrice, maxPrice, vendorID, query, sort)
}
