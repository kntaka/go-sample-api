package model

import (
	"time"
)

type User struct {
	ID        uint64    `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime;not null;DEFAULT:current_timestamp;"`
}
