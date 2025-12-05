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

func TestCustomerService_GetCustomerByID(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int64
	}

	type mockBehavior func(r *mocks.CustomerRepository)

	testCustomer := &model.Customer{
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
		want    *model.Customer
		wantErr bool
	}{
		{
			name: "Success - Customer Found",
			args: args{
				ctx: context.Background(),
				id:  42,
			},
			mock: func(r *mocks.CustomerRepository) {
				r.EXPECT().
					GetCustomerByID(mock.Anything, int64(42)).
					Return(testCustomer, nil)
			},
			want:    testCustomer,
			wantErr: false,
		},
		{
			name: "Error - Customer Not Found in Repository",
			args: args{
				ctx: context.Background(),
				id:  99,
			},
			mock: func(r *mocks.CustomerRepository) {
				repoErr := errors.New("sql: no rows in result set")
				r.EXPECT().
					GetCustomerByID(mock.Anything, int64(99)).
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
			mock: func(r *mocks.CustomerRepository) {
				repoErr := errors.New("connection failed")
				r.EXPECT().
					GetCustomerByID(mock.Anything, int64(10)).
					Return(nil, repoErr)
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewCustomerRepository(t)

			tt.mock(repo)

			s := service.NewCustomerService(repo)

			got, err := s.GetCustomerByID(tt.args.ctx, tt.args.id)

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
