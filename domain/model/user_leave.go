package model

import (
	"time"
)

type UserLeave struct {
	User      User      ``
	UserID    uint64    `json:"user_id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime;not null;DEFAULT:current_timestamp;"`
}
