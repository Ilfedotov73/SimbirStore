package dto

type ReviewDTO struct {
	ID         int64   `json:"id"`
	CustomerID int64   `json:"customerId"`
	ProductID  int64   `json:"productId"`
	Review     string  `json:"review"`
	Rating     float64 `json:"rating"`
}

type PagedReviews struct {
	Paging Paging       `json:"paging"`
	Items  []*ReviewDTO `json:"items"`
}
