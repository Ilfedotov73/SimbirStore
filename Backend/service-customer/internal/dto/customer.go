package dto

import "time"

type CustomerDTO struct {
	ID                 int64     `json:"id"`
	FirstName          string    `json:"firstName"`
	LastName           string    `json:"lastName"`
	PhoneNumber        string    `json:"phoneNumber"`
	PhotoURL           string    `json:"photoUrl"`
	CustomerTelegramID string    `json:"customerTelegramId"`
	CreateAt           time.Time `json:"createAt"`
	Login              string    `json:"login"`
	Email              string    `json:"email"`
}
