package service

import (
	"context"
	"fmt"
	"service-customer/internal/model"
)

//go:generate go run github.com/vektra/mockery/v2@latest --name=OfferRepository --output=../../mocks --outpkg=mocks --with-expecter=true
type OfferRepository interface {
	Create(ctx context.Context, offer *model.Offer) (*model.Offer, error)
}

type OfferService struct {
	repo OfferRepository
}

func NewOfferService(r OfferRepository) *OfferService {
	return &OfferService{repo: r}
}

func (s *OfferService) CreateOffer(ctx context.Context, offer model.Offer) (*model.Offer, error) {
	if offer.OfferPrice <= 0 {
		return nil, fmt.Errorf("offer price must be greater than zero")
	}

	if offer.Status == "" {
		offer.Status = model.OfferStatusPending
	}

	return s.repo.Create(ctx, &offer)
}
