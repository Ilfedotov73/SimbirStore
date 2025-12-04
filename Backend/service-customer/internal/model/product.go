package model

import "time"

type Product struct {
	ID              int64
	Name            string
	Price           float64
	PhotoURL        string
	CreateAt        time.Time
	Characteristics string
	ProductRating   float64
	VendorID        int64
}
