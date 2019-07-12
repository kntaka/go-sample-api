package repository

import (
	"github.com/takkee/go-sample-api/domain/model"
)

type UserActiveRepository interface {
	Store(u *model.UserActive) error
	Find(ua *model.UserActive) (*model.UserActive, error)
	Delete(ua *model.UserActive) (bool, error)
}
