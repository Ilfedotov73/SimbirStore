package model

import "time"

type UserRole int8

const (
	RoleCustomer UserRole = 0
	RoleVendor   UserRole = 1
)

type User struct {
	ID             int64
	FirstName      string
	LastName       string
	PhoneNumber    string
	PhotoURL       string
	UserTelegramID string
	CreateAt       time.Time
	Login          string
	Email          string
	Password       string
	Role           UserRole
}
