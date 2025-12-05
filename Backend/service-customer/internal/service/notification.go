package service

import (
	"context"
	"service-customer/internal/model"
	"time"
)

//go:generate go run github.com/vektra/mockery/v2@latest --name=NotificationRepository --output=../../mocks --outpkg=mocks --with-expecter=true
type NotificationRepository interface {
	GetByCustomerID(ctx context.Context, customerID int64, from, to time.Time, offset, limit int) ([]model.Notification, int, error)
}

type NotificationService struct {
	repo NotificationRepository
}

func NewNotificationService(r NotificationRepository) *NotificationService {
	return &NotificationService{repo: r}
}

func (s *NotificationService) GetNotifications(ctx context.Context, customerID int64, from, to time.Time, page, size int) ([]model.Notification, int, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 20
	}
	offset := (page - 1) * size
	return s.repo.GetByCustomerID(ctx, customerID, from, to, offset, size)
}
