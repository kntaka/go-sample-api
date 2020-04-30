package usecase_test

import (
	"context"
	"testing"

	"github.com/knwoop/go-sample-api/domain/model"
	"github.com/knwoop/go-sample-api/domain/service"
	"github.com/knwoop/go-sample-api/infrastructure/persistence/mock"
	"github.com/knwoop/go-sample-api/usecase"
)

func TestUserUseCase_Create(t *testing.T) {

	s := usecase.NewUserUseCase(
		service.NewUserService(mock.NewUserRepository(), mock.NewUserActiveRepository()),
		service.NewProfileService(mock.NewUserDetailRepository()))
	// normal test
	t.Run("UserID=1", func(t *testing.T) {
		t.Log("UserID=1")
		ud := &model.UserDetail{}
		ud.Name = "test user"
		ud.Email = "sample@test.com"
		err := s.Create(context.TODO(), ud)
		if err != nil {
			actual := err
			t.Errorf("got: %v\nwant: %v", actual, nil)
		}
	})

	// email empty error
	t.Run("email=empty", func(t *testing.T) {
		t.Log("email=empty")
		ud := &model.UserDetail{}
		ud.Name = "test user"
		err := s.Create(context.TODO(), ud)
		if err == nil {
			t.Errorf("got: %v\nwant: %v", nil, "Error")
		}
	})

	// name error
	t.Run("name=empty", func(t *testing.T) {
		t.Log("name=empty")
		ud := &model.UserDetail{}
		ud.Email = "sample@test.com"
		err := s.Create(context.TODO(), ud)
		if err == nil {
			t.Errorf("got: %v\nwant: %v", nil, "Error")
		}
	})
}
