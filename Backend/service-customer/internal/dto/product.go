package dto

import "time"

type ProductSummary struct {
	ID              int64   `json:"id"`
	Name            string  `json:"name"`
	Price           float64 `json:"price"`
	PhotoURL        string  `json:"photoUrl"`
	Characteristics string  `json:"characteristics"`
	ProductRating   float64 `json:"productRating"`
	VendorID        int64   `json:"vendorId"`
}

type ProductDetail struct {
	ProductSummary
	CreateAt time.Time `json:"createAt"`
}

type FullProductResponse struct {
	Product       *ProductDetail `json:"product"`
	Vendor        *VendorDTO     `json:"vendor"`
	ReviewsCount  int            `json:"reviewsCount"`
	AverageRating float64        `json:"aveargeRating"`
}

type PagedProducts struct {
	Paging Paging            `json:"paging"`
	Items  []*ProductSummary `json:"items"`
}
