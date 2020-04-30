package datastore

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/knwoop/go-sample-api/domain/model"
	"github.com/knwoop/go-sample-api/domain/repository"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db}
}

func (ur *userRepository) Store(u *model.User) error {
	return ur.db.Save(u).Error
}

func (ur *userRepository) Find(userID uint64) (*model.User, error) {

	u := &model.User{ID: userID}
	err := ur.db.Find(u).Error
	if err != nil {
		return nil, fmt.Errorf("sql error", err)
	}

	return u, nil
}
