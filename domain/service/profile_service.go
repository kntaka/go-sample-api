package service

import (
	"context"
	"errors"
	"github.com/takkee/go-sample-api/domain/model"
	"github.com/takkee/go-sample-api/domain/repository"
)

// creating user service
type ProfileService interface {
	Create(ctx context.Context, ud *model.UserDetail) error
}

func NewProfileService(udr repository.UserDetailRepository) ProfileService {
	return &profileService{udr: udr}
}

type profileService struct {
	udr repository.UserDetailRepository
}

func (s *profileService) Create(ctx context.Context, ud *model.UserDetail) error {
	if ud.UserID == 0 {
		return errors.New("need to set user id before calling this method")
	}
	// check user id
	if ud.Email == "" {
		return errors.New("need to set email before calling this method")
	}

	// check user id
	if ud.Name == "" {
		return errors.New("need to set name before calling this method")
	}
	// create user detail model
	uErr := s.udr.Store(ud)
	if uErr != nil {
		return uErr
	}

	return nil
}
