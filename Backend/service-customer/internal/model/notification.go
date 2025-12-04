package model

import "time"

type Notification struct {
	ID       int64
	Text     string
	EntityID int64
	CreateAt time.Time
}
