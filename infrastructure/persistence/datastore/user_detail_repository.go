package datastore

import (
	"github.com/jinzhu/gorm"
	"github.com/takkee/go-sample-api/domain/model"
	"github.com/takkee/go-sample-api/domain/repository"
)

type userDetailRepository struct {
	db *gorm.DB
}

func NewUserDetailRepository(db *gorm.DB) repository.UserDetailRepository {
	return &userDetailRepository{db}
}

func (ur *userDetailRepository) Store(u *model.UserDetail) error {
	return ur.db.Save(u).Error
}

func (ur *userDetailRepository) FindOneByUserID(userID uint64) (*model.UserDetail, error) {
	return nil, nil
}

