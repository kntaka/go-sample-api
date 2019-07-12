package service

import (
	"context"
	"github.com/takkee/go-sample-api/domain/model"
	"github.com/takkee/go-sample-api/domain/repository"
)

// creating user service
type UserService interface {
	Create(ctx context.Context) (uint64, error)
}

func NewUserService(ur repository.UserRepository, uar repository.UserActiveRepository) UserService {
	return &userService{ur: ur, uar: uar}
}

type userService struct {
	ur repository.UserRepository
	uar repository.UserActiveRepository
}

func (s *userService) Create(ctx context.Context) (uint64, error) {
	// create user model
	u    := &model.User{}
	uErr := s.ur.Store(u)
	if uErr != nil {
		return 0, uErr
	}

	// create user status
	userId := u.ID
	ua     := &model.UserActive{}
	ua.UserID = userId
	uaErr := s.uar.Store(ua)
	if uaErr != nil {
		return 0, uaErr
	}

	return u.ID, nil
}
