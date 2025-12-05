package model

import "time"

type OfferStatus string

const (
	OfferStatusPending  OfferStatus = "pending"
	OfferStatusAccepted OfferStatus = "accepted"
	OfferStatusRejected OfferStatus = "rejected"
	OfferStatusCounter  OfferStatus = "counter"
)

type Offer struct {
	ID         int64       `db:"id"`
	ProductID  int64       `db:"product_id"`
	VendorID   int64       `db:"vendor_id"`
	BuyerID    int64       `db:"buyer_id"`
	OfferPrice float64     `db:"offer_price"`
	Message    string      `db:"message"`
	Status     OfferStatus `db:"status"`
	CreateAt   time.Time   `db:"create_at"`
}
