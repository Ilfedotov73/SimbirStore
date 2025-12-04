package model

import "time"

type Vendor struct {
	ID               int64
	FirstName        string
	LastName         string
	PhoneNumber      string
	PhotoURL         string
	VendorTelegramID string
	CreateAt         time.Time
	Login            string
	Email            string
	Password         string
}
