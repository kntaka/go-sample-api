package service_test

import (
	"context"
	"github.com/kntaka/go-sample-api/domain/service"
	"github.com/kntaka/go-sample-api/infrastructure/persistence/mock"
	"testing"
)

func TestUserService_Create(t *testing.T) {

	s := service.NewUserService(mock.NewUserRepository(), mock.NewUserActiveRepository())
	// normal test
	t.Run("create test 1", func(t *testing.T) {
		t.Log("create test 1")
		_, err := s.Create(context.TODO())
		if err != nil {
			actual := "Error"
			t.Errorf("got: %v\nwant: %v", actual, nil)
		}
	})

	t.Run("create test 2", func(t *testing.T) {
		t.Log("create test 2")
		_, err := s.Create(context.TODO())
		if err != nil {
			actual := "Error"
			t.Errorf("got: %v\nwant: %v", actual, nil)
		}
	})
}
