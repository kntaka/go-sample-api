package mock

import (
	"time"

	"github.com/knwoop/go-sample-api/domain/model"
	"github.com/knwoop/go-sample-api/domain/repository"
)

type userDetailRepository struct {
}

func NewUserDetailRepository() repository.UserDetailRepository {
	return &userDetailRepository{}
}

func (ur *userDetailRepository) Store(u *model.UserDetail) error {
	return nil
}

func (ur *userDetailRepository) FindOneByUserID(userID uint64) (*model.UserDetail, error) {

	ud := &model.UserDetail{ID: userID}
	ud.CreatedAt = time.Time{}
	return ud, nil
}
