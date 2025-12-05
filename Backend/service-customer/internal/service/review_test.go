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

func TestReviewService_GetReviewStats(t *testing.T) {
	type args struct {
		ctx       context.Context
		productID int64
	}

	type mockBehavior func(r *mocks.ReviewRepository)

	tests := []struct {
		name      string
		args      args
		mock      mockBehavior
		wantCount int
		wantAvg   float64
		wantErr   bool
	}{
		{
			name: "Success - Valid Stats",
			args: args{
				ctx:       context.Background(),
				productID: 101,
			},
			mock: func(r *mocks.ReviewRepository) {
				r.EXPECT().
					GetStats(mock.Anything, int64(101)).
					Return(50, 4.2, nil)
			},
			wantCount: 50,
			wantAvg:   4.2,
			wantErr:   false,
		},
		{
			name: "Success - No Reviews",
			args: args{
				ctx:       context.Background(),
				productID: 102,
			},
			mock: func(r *mocks.ReviewRepository) {
				r.EXPECT().
					GetStats(mock.Anything, int64(102)).
					Return(0, 0.0, nil)
			},
			wantCount: 0,
			wantAvg:   0.0,
			wantErr:   false,
		},
		{
			name: "Repository Error",
			args: args{
				ctx:       context.Background(),
				productID: 103,
			},
			mock: func(r *mocks.ReviewRepository) {
				r.EXPECT().
					GetStats(mock.Anything, int64(103)).
					Return(0, 0.0, errors.New("db error on stats"))
			},
			wantCount: 0,
			wantAvg:   0.0,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewReviewRepository(t)
			tt.mock(repo)

			s := service.NewReviewService(repo)

			gotCount, gotAvg, err := s.GetReviewStats(tt.args.ctx, tt.args.productID)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantCount, gotCount)
				assert.Equal(t, tt.wantAvg, gotAvg)
			}
		})
	}
}

func TestReviewService_GetReviews(t *testing.T) {
	type args struct {
		ctx       context.Context
		productID int64
		page      int
		size      int
	}

	type mockBehavior func(r *mocks.ReviewRepository)

	reviewList := []model.Review{
		{ID: 1, Review: "Great product", Rating: 5.0},
		{ID: 2, Review: "Meh", Rating: 3.0},
	}

	tests := []struct {
		name       string
		args       args
		mock       mockBehavior
		wantResult []model.Review
		wantTotal  int
		wantErr    bool
	}{
		{
			name: "Success - First Page",
			args: args{
				ctx:       context.Background(),
				productID: 201,
				page:      1,
				size:      2,
			},
			mock: func(r *mocks.ReviewRepository) {
				r.EXPECT().
					GetByProductID(mock.Anything, int64(201), 0, 2).
					Return(reviewList, 10, nil)
			},
			wantResult: reviewList,
			wantTotal:  10,
			wantErr:    false,
		},
		{
			name: "Success - Second Page",
			args: args{
				ctx:       context.Background(),
				productID: 202,
				page:      3,
				size:      5,
			},
			mock: func(r *mocks.ReviewRepository) {
				r.EXPECT().
					GetByProductID(mock.Anything, int64(202), 10, 5).
					Return([]model.Review{}, 25, nil)
			},
			wantResult: []model.Review{},
			wantTotal:  25,
			wantErr:    false,
		},
		{
			name: "Repository Error",
			args: args{
				ctx:       context.Background(),
				productID: 203,
				page:      1,
				size:      10,
			},
			mock: func(r *mocks.ReviewRepository) {
				r.EXPECT().
					GetByProductID(mock.Anything, int64(203), 0, 10).
					Return(nil, 0, errors.New("failed to fetch reviews"))
			},
			wantResult: nil,
			wantTotal:  0,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewReviewRepository(t)
			tt.mock(repo)

			s := service.NewReviewService(repo)

			got, total, err := s.GetReviews(
				tt.args.ctx,
				tt.args.productID,
				tt.args.page,
				tt.args.size,
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
