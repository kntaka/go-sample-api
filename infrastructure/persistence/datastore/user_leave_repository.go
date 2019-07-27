package datastore

import (
	"github.com/jinzhu/gorm"
	"github.com/kntaka/go-sample-api/domain/model"
	"github.com/kntaka/go-sample-api/domain/repository"
)

type UserLeaveRepository interface {
	Store(user *model.UserLeave) error
	FindOneByUserId(userID uint64) (*model.UserLeave, error)
}

type userLeaveRepository struct {
	db *gorm.DB
}

func NewUserLeaveRepository(db *gorm.DB) repository.UserLeaveRepository {
	return &userLeaveRepository{db}
}

func (ur *userLeaveRepository) Store(ul *model.UserLeave) error {
	return ur.db.Save(ul).Error
}

func (ur *userLeaveRepository) FindOneByUserId(userID uint64) (*model.UserLeave, error) {
	return nil, nil
}

