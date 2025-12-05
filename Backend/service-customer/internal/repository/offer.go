package repository

import (
	"context"
	"database/sql"
	"service-customer/internal/model"
)

type OfferRepository struct {
	db *sql.DB
}

func NewOfferRepository(db *sql.DB) *OfferRepository {
	return &OfferRepository{db: db}
}

func (r *OfferRepository) Create(ctx context.Context, offer *model.Offer) (*model.Offer, error) {
	query := `
		INSERT INTO offers (product_id, buyer_id, vendor_id, offer_price, message, status, create_at)
		VALUES ($1, $2, $3, $4, $5, $6, NOW())
		RETURNING id, create_at, status
	`
	err := r.db.QueryRowContext(
		ctx,
		query,
		offer.ProductID,
		offer.BuyerID,
		offer.VendorID,
		offer.OfferPrice,
		offer.Message,
		offer.Status,
	).Scan(&offer.ID, &offer.CreateAt, &offer.Status)

	if err != nil {
		return nil, err
	}
	return offer, nil
}
