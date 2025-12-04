package model

type Review struct {
	ID         int64
	CustomerID int64
	ProductID  int64
	Review     string
	Rating     float64
}
