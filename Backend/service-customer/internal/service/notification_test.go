package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"service-customer/internal/model"
	"service-customer/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNotificationService_GetNotifications(t *testing.T) {
	timeFrom := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	timeTo := time.Date(2025, 1, 31, 23, 59, 59, 0, time.UTC)

	notifications := []model.Notification{
		{ID: 1, Text: "Offer accepted", CreateAt: timeFrom.Add(time.Hour)},
		{ID: 2, Text: "New review", CreateAt: timeFrom.Add(2 * time.Hour)},
	}

	type args struct {
		ctx        context.Context
		customerID int64
		from       time.Time
		to         time.Time
		page       int
		size       int
	}

	type mockBehavior func(r *mocks.NotificationRepository)

	tests := []struct {
		name       string
		args       args
		mock       mockBehavior
		wantResult []model.Notification
		wantTotal  int
		wantErr    bool
	}{
		{
			name: "Success - Valid Pagination and Time Range",
			args: args{
				ctx:        context.Background(),
				customerID: 10,
				from:       timeFrom,
				to:         timeTo,
				page:       2,
				size:       5,
			},
			mock: func(r *mocks.NotificationRepository) {
				r.EXPECT().
					GetByCustomerID(mock.Anything, int64(10), timeFrom, timeTo, 5, 5).
					Return(notifications, 15, nil)
			},
			wantResult: notifications,
			wantTotal:  15,
			wantErr:    false,
		},
		{
			name: "Default Pagination (page < 1, size < 1)",
			args: args{
				ctx:        context.Background(),
				customerID: 20,
				from:       timeFrom,
				to:         timeTo,
				page:       0,
				size:       0,
			},
			mock: func(r *mocks.NotificationRepository) {
				r.EXPECT().
					GetByCustomerID(mock.Anything, int64(20), timeFrom, timeTo, 0, 20).
					Return([]model.Notification{}, 5, nil)
			},
			wantResult: []model.Notification{},
			wantTotal:  5,
			wantErr:    false,
		},
		{
			name: "Repository Error",
			args: args{
				ctx:        context.Background(),
				customerID: 30,
				from:       timeFrom,
				to:         timeTo,
				page:       1,
				size:       10,
			},
			mock: func(r *mocks.NotificationRepository) {
				r.EXPECT().
					GetByCustomerID(mock.Anything, int64(30), timeFrom, timeTo, 0, 10).
					Return(nil, 0, errors.New("database connection failed"))
			},
			wantResult: nil,
			wantTotal:  0,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewNotificationRepository(t)
			tt.mock(repo)

			s := NewNotificationService(repo)

			got, total, err := s.GetNotifications(
				tt.args.ctx,
				tt.args.customerID,
				tt.args.from,
				tt.args.to,
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
