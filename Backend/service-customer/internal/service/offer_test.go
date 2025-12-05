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

func TestOfferService_CreateOffer(t *testing.T) {
	createdOffer := &model.Offer{
		ID:         1,
		ProductID:  100,
		VendorID:   50,
		BuyerID:    5,
		OfferPrice: 1500.00,
		Status:     model.OfferStatusPending,
		CreateAt:   time.Now(),
	}

	type args struct {
		ctx   context.Context
		offer model.Offer
	}

	type mockBehavior func(r *mocks.OfferRepository)

	tests := []struct {
		name    string
		args    args
		mock    mockBehavior
		want    *model.Offer
		wantErr bool
	}{
		{
			name: "Success - Default Status (Pending)",
			args: args{
				ctx: context.Background(),
				offer: model.Offer{
					ProductID:  100,
					VendorID:   50,
					BuyerID:    5,
					OfferPrice: 1500.00,
					Status:     "",
				},
			},
			mock: func(r *mocks.OfferRepository) {
				matcher := mock.MatchedBy(func(o *model.Offer) bool {
					return o.Status == model.OfferStatusPending && o.OfferPrice == 1500.00
				})
				r.EXPECT().
					Create(mock.Anything, matcher).
					Return(createdOffer, nil)
			},
			want:    createdOffer,
			wantErr: false,
		},
		{
			name: "Success - Custom Status (Counter)",
			args: args{
				ctx: context.Background(),
				offer: model.Offer{
					ProductID:  101,
					VendorID:   51,
					BuyerID:    6,
					OfferPrice: 200.00,
					Status:     model.OfferStatusCounter,
				},
			},
			mock: func(r *mocks.OfferRepository) {
				matcher := mock.MatchedBy(func(o *model.Offer) bool {
					return o.Status == model.OfferStatusCounter && o.OfferPrice == 200.00
				})
				customOffer := &model.Offer{ID: 2, Status: model.OfferStatusCounter}
				r.EXPECT().
					Create(mock.Anything, matcher).
					Return(customOffer, nil)
			},
			want:    &model.Offer{ID: 2, Status: model.OfferStatusCounter},
			wantErr: false,
		},
		{
			name: "Failure - Zero Price Validation",
			args: args{
				ctx: context.Background(),
				offer: model.Offer{
					OfferPrice: 0.00,
				},
			},
			mock:    func(r *mocks.OfferRepository) {},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Failure - Negative Price Validation",
			args: args{
				ctx: context.Background(),
				offer: model.Offer{
					OfferPrice: -10.00,
				},
			},
			mock:    func(r *mocks.OfferRepository) {},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Failure - Repository Error",
			args: args{
				ctx: context.Background(),
				offer: model.Offer{
					OfferPrice: 1000.00,
				},
			},
			mock: func(r *mocks.OfferRepository) {
				r.EXPECT().
					Create(mock.Anything, mock.AnythingOfType("*model.Offer")).
					Return(nil, errors.New("database connection failed"))
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewOfferRepository(t)
			tt.mock(repo)

			s := service.NewOfferService(repo)

			got, err := s.CreateOffer(tt.args.ctx, tt.args.offer)

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
