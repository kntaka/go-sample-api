package service_test

import (
	"context"
	"testing"

	"github.com/knwoop/go-sample-api/domain/model"
	"github.com/knwoop/go-sample-api/domain/service"
	"github.com/knwoop/go-sample-api/infrastructure/persistence/mock"
)

func TestProfileService_Create(t *testing.T) {
	var userID uint64 = 100

	s := service.NewProfileService(mock.NewUserDetailRepository())
	// normal test
	t.Run("UserID=1", func(t *testing.T) {
		t.Log("UserID=1")
		ud := &model.UserDetail{}
		ud.UserID = userID
		ud.Name = "test user"
		ud.Email = "sample@test.com"
		err := s.Create(context.TODO(), ud)
		if err != nil {
			actual := "Error"
			t.Errorf("got: %v\nwant: %v", actual, nil)
		}
	})

	// user id empty error
	t.Run("UserID=empty", func(t *testing.T) {
		t.Log("UserID=empty")
		ud := &model.UserDetail{}
		ud.Name = "test user"
		ud.Email = "sample@test.com"
		err := s.Create(context.TODO(), ud)
		if err == nil {
			t.Errorf("got: %v\nwant: %v", nil, "Error")
		}
	})

	// email empty error
	t.Run("email=empty", func(t *testing.T) {
		t.Log("email=empty")
		ud := &model.UserDetail{}
		ud.UserID = userID
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
		ud.UserID = userID
		ud.Email = "sample@test.com"
		err := s.Create(context.TODO(), ud)
		if err == nil {
			t.Errorf("got: %v\nwant: %v", nil, "Error")
		}
	})
}
