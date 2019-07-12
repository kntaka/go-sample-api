package model

import (
	"time"
)

type UserDetail struct {
	ID         uint64    `gorm:"primary_key"`
	User       User      ``
	UserID     uint64    `json:"user_id" gorm:"index"`
	Name       string    `json:"name" gorm:"not null;size:255"`
	Email      string    `json:"birthday_at" gorm:"not null;size:255"`
	CreatedAt  time.Time `json:"created_at" gorm:"type:datetime;not null;DEFAULT:current_timestamp;"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"type:datetime;not null;DEFAULT:current_timestamp;"`
}
