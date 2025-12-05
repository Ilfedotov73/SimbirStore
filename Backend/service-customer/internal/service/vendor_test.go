package service_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"service-customer/internal/model"
	"service-customer/internal/service"
	"service-customer/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestVendorService_GetVendorByID(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int64
	}

	type mockBehavior func(r *mocks.VendorRepository)

	testVendor := &model.Vendor{
		ID:        42,
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "jane@example.com",
		CreateAt:  time.Now(),
	}

	tests := []struct {
		name    string
		args    args
		mock    mockBehavior
		want    *model.Vendor
		wantErr bool
	}{
		{
			name: "Success - Vendor Found",
			args: args{
				ctx: context.Background(),
				id:  42,
			},
			mock: func(r *mocks.VendorRepository) {
				r.EXPECT().
					GetVendorByID(mock.Anything, int64(42)).
					Return(testVendor, nil)
			},
			want:    testVendor,
			wantErr: false,
		},
		{
			name: "Error - Vendor Not Found in Repository",
			args: args{
				ctx: context.Background(),
				id:  99,
			},
			mock: func(r *mocks.VendorRepository) {
				repoErr := errors.New("sql: no rows in result set")
				r.EXPECT().
					GetVendorByID(mock.Anything, int64(99)).
					Return(nil, repoErr)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Error - General DB Error",
			args: args{
				ctx: context.Background(),
				id:  10,
			},
			mock: func(r *mocks.VendorRepository) {
				repoErr := errors.New("connection failed")
				r.EXPECT().
					GetVendorByID(mock.Anything, int64(10)).
					Return(nil, repoErr)
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewVendorRepository(t)

			tt.mock(repo)

			s := service.NewVendorService(repo)

			got, err := s.GetVendorByID(tt.args.ctx, tt.args.id)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
