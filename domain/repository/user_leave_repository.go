package repository

import (
	"github.com/kntaka/go-sample-api/domain/model"
)

type UserLeaveRepository interface {
	Store(u *model.UserLeave) error
	FindOneByUserId(userID uint64) (*model.UserLeave, error)
}
