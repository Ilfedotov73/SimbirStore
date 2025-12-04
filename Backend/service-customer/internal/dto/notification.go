package dto

import "time"

type NotificationDTO struct {
	ID       int64     `json:"id"`
	Text     string    `json:"text"`
	EntityID int64     `json:"entityId"`
	CreateAt time.Time `json:"createAt"`
}

type PagedNotifications struct {
	Paging Paging             `json:"paging"`
	Items  []*NotificationDTO `json:"items"`
}
