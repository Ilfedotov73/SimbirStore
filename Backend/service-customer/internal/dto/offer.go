package dto

import "time"

type OfferRequest struct {
	BuyerID    int64   `json:"buyerId" binding:"required"`
	VendorID   int64   `json:"vendorId" binding:"required"`
	OfferPrice float64 `json:"offerPrice" binding:"required"`
	Message    string  `json:"message"`
}

type OfferResponse struct {
	ID         int64     `json:"id"`
	ProductID  int64     `json:"productId"`
	BuyerID    int64     `json:"buyerId"`
	VendorID   int64     `json:"vendorId"`
	OfferPrice float64   `json:"offerPrice"`
	Message    string    `json:"message"`
	Status     string    `json:"status"`
	CreateAt   time.Time `json:"createAt"`
}
