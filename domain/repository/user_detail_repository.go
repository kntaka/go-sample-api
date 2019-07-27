package repository

import (
	"github.com/kntaka/go-sample-api/domain/model"
)

type UserDetailRepository interface {
	Store(u *model.UserDetail) error
	FindOneByUserID(userID uint64) (*model.UserDetail, error)
}
