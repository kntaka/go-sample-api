package mock

import (
	"github.com/takkee/go-sample-api/domain/model"
	"github.com/takkee/go-sample-api/domain/repository"
	"time"
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
