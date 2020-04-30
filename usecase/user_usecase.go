package usecase

import (
	"context"

	"github.com/knwoop/go-sample-api/domain/model"
	"github.com/knwoop/go-sample-api/domain/service"
)

type UserUseCase interface {
	Create(ctx context.Context, ud *model.UserDetail) error
}

func NewUserUseCase(
	us service.UserService,
	ps service.ProfileService,
) UserUseCase {
	return &userUseCase{us, ps}
}

type userUseCase struct {
	us service.UserService
	ps service.ProfileService
}

// create a user
func (uc *userUseCase) Create(ctx context.Context, ud *model.UserDetail) error {
	// create a user
	userID, usErr := uc.us.Create(ctx)
	if usErr != nil {
		return usErr
	}

	// create a user
	ud.UserID = userID
	psErr := uc.ps.Create(ctx, ud)
	if psErr != nil {
		return psErr
	}

	return nil
}
