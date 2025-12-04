package service

import (
	"context"
	"service-customer/internal/model"
)

type ReviewRepository interface {
	GetStats(ctx context.Context, productID int64) (int, float64, error)
	GetByProductID(ctx context.Context, productID int64, offset, limit int) ([]model.Review, int, error)
}

type ReviewService struct {
	repo ReviewRepository
}

func NewReviewService(r ReviewRepository) *ReviewService {
	return &ReviewService{repo: r}
}

func (s *ReviewService) GetReviewStats(ctx context.Context, productID int64) (int, float64, error) {
	return s.repo.GetStats(ctx, productID)
}

func (s *ReviewService) GetReviews(ctx context.Context, productID int64, page, size int) ([]model.Review, int, error) {
	offset := (page - 1) * size
	return s.repo.GetByProductID(ctx, productID, offset, size)
}
