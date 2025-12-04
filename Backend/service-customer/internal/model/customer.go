package model

import "time"

type Customer struct {
	ID                 int64
	FirstName          string
	LastName           string
	PhoneNumber        string
	PhotoURL           string
	CustomerTelegramID string
	CreateAt           time.Time
	Login              string
	Email              string
	Password           string
}
