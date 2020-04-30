package datastore

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/knwoop/go-sample-api/domain/model"
	"github.com/knwoop/go-sample-api/domain/repository"
)

type userActiveRepository struct {
	db *gorm.DB
}

func NewUserActiveRepository(db *gorm.DB) repository.UserActiveRepository {
	return &userActiveRepository{db}
}

func (ur *userActiveRepository) Store(u *model.UserActive) error {
	return ur.db.Save(u).Error
}

func (ur *userActiveRepository) Find(u *model.UserActive) (*model.UserActive, error) {

	err := ur.db.Find(&u).Error
	if err != nil {
		return nil, fmt.Errorf("sql error", err)
	}

	return u, nil
}

func (ur *userActiveRepository) Delete(u *model.UserActive) (bool, error) {

	err := ur.db.Find(&u).Error
	if err != nil {
		return false, fmt.Errorf("sql error", err)
	}

	return true, nil
}
