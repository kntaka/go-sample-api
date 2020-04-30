package controller

import (
	"context"

	"github.com/knwoop/go-sample-api/domain/model"
	"github.com/knwoop/go-sample-api/infrastructure/api/validator"
	"github.com/knwoop/go-sample-api/usecase"
)

type UserController interface {
	Get(ctx context.Context) error
	Create(ctx context.Context, req *validator.UserCreateRequest) error
	Update(ctx context.Context) error
	Delete(ctx context.Context) error
}

func NewUserController(uc usecase.UserUseCase) UserController {
	return &userController{uc}
}

type userController struct {
	uc usecase.UserUseCase
}

func (uc *userController) Get(ctx context.Context) error {
	// 未実装
	return nil
}

func (uc *userController) Create(ctx context.Context, req *validator.UserCreateRequest) error {

	ud := &model.UserDetail{}
	ud.Name = req.Name
	ud.Email = req.Email

	err := uc.uc.Create(ctx, ud)
	if err != nil {
		return err
	}
	return nil
}

func (uc *userController) Update(ctx context.Context) error {
	// 未実装
	return nil
}

func (uc *userController) Delete(ctx context.Context) error {
	// 未実装
	return nil
}
