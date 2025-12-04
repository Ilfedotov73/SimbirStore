package dto

import "time"

type VendorDTO struct {
	ID               int64     `json:"id"`
	FirstName        string    `json:"firstName"`
	LastName         string    `json:"lastName"`
	PhoneNumber      string    `json:"phoneNumber"`
	PhotoURL         string    `json:"photoUrl"`
	VendorTelegramID string    `json:"vendorTelegramId"`
	CreateAt         time.Time `json:"createAt"`
}
