package service_test

import (
	"context"
	"errors"
	"testing"

	"service-customer/internal/model"
	"service-customer/internal/service"
	"service-customer/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestProductService_GetProductByID(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int64
	}

	type mockBehavior func(r *mocks.ProductRepository)

	tests := []struct {
		name    string
		args    args
		mock    mockBehavior
		want    *model.Product
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			mock: func(r *mocks.ProductRepository) {
				r.EXPECT().
					GetProductByID(mock.Anything, int64(1)).
					Return(&model.Product{ID: 1, Name: "Test Product"}, nil)
			},
			want:    &model.Product{ID: 1, Name: "Test Product"},
			wantErr: false,
		},
		{
			name: "Repository Error",
			args: args{
				ctx: context.Background(),
				id:  2,
			},
			mock: func(r *mocks.ProductRepository) {
				r.EXPECT().
					GetProductByID(mock.Anything, int64(2)).
					Return(nil, errors.New("db error"))
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewProductRepository(t)
			vendorService := mocks.NewProductVendorService(t)
			reviewService := mocks.NewProductReviewService(t)

			tt.mock(repo)

			s := service.NewProductService(repo, vendorService, reviewService)

			got, err := s.GetProductByID(tt.args.ctx, tt.args.id)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestProductService_GetFullProductData(t *testing.T) {
	type args struct {
		ctx       context.Context
		productID int64
	}

	type mockBehavior func(r *mocks.ProductRepository, v *mocks.ProductVendorService, rev *mocks.ProductReviewService)

	testProduct := &model.Product{ID: 10, VendorID: 5, Name: "Phone"}
	testVendor := &model.Vendor{ID: 5, FirstName: "John"}

	tests := []struct {
		name       string
		args       args
		mock       mockBehavior
		wantProd   *model.Product
		wantVendor *model.Vendor
		wantCount  int
		wantAvg    float64
		wantErr    bool
	}{
		{
			name: "Success with all data",
			args: args{
				ctx:       context.Background(),
				productID: 10,
			},
			mock: func(r *mocks.ProductRepository, v *mocks.ProductVendorService, rev *mocks.ProductReviewService) {
				r.EXPECT().GetProductByID(mock.Anything, int64(10)).Return(testProduct, nil)
				v.EXPECT().GetVendorByID(mock.Anything, int64(5)).Return(testVendor, nil)
				rev.EXPECT().GetReviewStats(mock.Anything, int64(10)).Return(15, 4.5, nil)
			},
			wantProd:   testProduct,
			wantVendor: testVendor,
			wantCount:  15,
			wantAvg:    4.5,
			wantErr:    false,
		},
		{
			name: "Product Not Found",
			args: args{
				ctx:       context.Background(),
				productID: 999,
			},
			mock: func(r *mocks.ProductRepository, v *mocks.ProductVendorService, rev *mocks.ProductReviewService) {
				r.EXPECT().GetProductByID(mock.Anything, int64(999)).Return(nil, errors.New("not found"))
			},
			wantProd:   nil,
			wantVendor: nil,
			wantCount:  0,
			wantAvg:    0,
			wantErr:    true,
		},
		{
			name: "Vendor Not Found",
			args: args{
				ctx:       context.Background(),
				productID: 10,
			},
			mock: func(r *mocks.ProductRepository, v *mocks.ProductVendorService, rev *mocks.ProductReviewService) {
				r.EXPECT().GetProductByID(mock.Anything, int64(10)).Return(testProduct, nil)
				v.EXPECT().GetVendorByID(mock.Anything, int64(5)).Return(nil, errors.New("vendor unavailable"))
			},
			wantProd:   nil,
			wantVendor: nil,
			wantCount:  0,
			wantAvg:    0,
			wantErr:    true,
		},
		{
			name: "Review Service Error (Should be ignored)",
			args: args{
				ctx:       context.Background(),
				productID: 10,
			},
			mock: func(r *mocks.ProductRepository, v *mocks.ProductVendorService, rev *mocks.ProductReviewService) {
				r.EXPECT().GetProductByID(mock.Anything, int64(10)).Return(testProduct, nil)
				v.EXPECT().GetVendorByID(mock.Anything, int64(5)).Return(testVendor, nil)
				rev.EXPECT().GetReviewStats(mock.Anything, int64(10)).Return(0, 0.0, errors.New("review service down"))
			},
			wantProd:   testProduct,
			wantVendor: testVendor,
			wantCount:  0,
			wantAvg:    0.0,
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewProductRepository(t)
			vendorService := mocks.NewProductVendorService(t)
			reviewService := mocks.NewProductReviewService(t)

			tt.mock(repo, vendorService, reviewService)

			s := service.NewProductService(repo, vendorService, reviewService)

			gotProd, gotVend, gotCount, gotAvg, err := s.GetFullProductData(tt.args.ctx, tt.args.productID)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantProd, gotProd)
				assert.Equal(t, tt.wantVendor, gotVend)
				assert.Equal(t, tt.wantCount, gotCount)
				assert.Equal(t, tt.wantAvg, gotAvg)
			}
		})
	}
}

func TestProductService_GetProducts(t *testing.T) {
	type args struct {
		ctx      context.Context
		page     int
		size     int
		minPrice float64
		maxPrice float64
		vendorID int64
		query    string
		sort     string
	}

	type mockBehavior func(r *mocks.ProductRepository)

	tests := []struct {
		name       string
		args       args
		mock       mockBehavior
		wantResult []model.Product
		wantTotal  int
		wantErr    bool
	}{
		{
			name: "Default Pagination (page < 1, size < 1)",
			args: args{
				ctx:      context.Background(),
				page:     0,
				size:     0,
				vendorID: 0,
			},
			mock: func(r *mocks.ProductRepository) {
				r.EXPECT().
					GetProducts(mock.Anything, 0, 20, 0.0, 0.0, int64(0), "", "").
					Return([]model.Product{{ID: 1}}, 1, nil)
			},
			wantResult: []model.Product{{ID: 1}},
			wantTotal:  1,
			wantErr:    false,
		},
		{
			name: "Custom Pagination (page 2, size 10)",
			args: args{
				ctx:      context.Background(),
				page:     2,
				size:     10,
				minPrice: 100,
				vendorID: 55,
			},
			mock: func(r *mocks.ProductRepository) {
				r.EXPECT().
					GetProducts(mock.Anything, 10, 10, 100.0, 0.0, int64(55), "", "").
					Return([]model.Product{}, 0, nil)
			},
			wantResult: []model.Product{},
			wantTotal:  0,
			wantErr:    false,
		},
		{
			name: "Repo Error",
			args: args{
				ctx:  context.Background(),
				page: 1,
				size: 10,
			},
			mock: func(r *mocks.ProductRepository) {
				r.EXPECT().
					GetProducts(mock.Anything, 0, 10, 0.0, 0.0, int64(0), "", "").
					Return(nil, 0, errors.New("failed to fetch"))
			},
			wantResult: nil,
			wantTotal:  0,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewProductRepository(t)
			vendorService := mocks.NewProductVendorService(t)
			reviewService := mocks.NewProductReviewService(t)

			tt.mock(repo)

			s := service.NewProductService(repo, vendorService, reviewService)

			got, total, err := s.GetProducts(
				tt.args.ctx,
				tt.args.page,
				tt.args.size,
				tt.args.minPrice,
				tt.args.maxPrice,
				tt.args.vendorID,
				tt.args.query,
				tt.args.sort,
			)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantResult, got)
				assert.Equal(t, tt.wantTotal, total)
			}
		})
	}
}
