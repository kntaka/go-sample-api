package mock

import (
	"github.com/knwoop/go-sample-api/domain/model"
	"github.com/knwoop/go-sample-api/domain/repository"
)

type userActiveRepository struct {
}

func NewUserActiveRepository() repository.UserActiveRepository {
	return &userActiveRepository{}
}

func (ur *userActiveRepository) Store(u *model.UserActive) error {
	return nil
}

func (ur *userActiveRepository) Find(u *model.UserActive) (*model.UserActive, error) {
	return u, nil
}

func (ur *userActiveRepository) Delete(u *model.UserActive) (bool, error) {
	return true, nil
}
