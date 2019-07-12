package handler

import (
	"context"
	"github.com/takkee/go-sample-api/domain/model"
	"github.com/takkee/go-sample-api/infrastructure/api/validator"
	"github.com/takkee/go-sample-api/interface/controller"

	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler interface {
	CreateUser(c *gin.Context)
}

func NewUserHandler(c controller.UserController) UserHandler {
	return &userHandler{userController: c}
}

type userHandler struct {
	userController controller.UserController
}

func (sh *userHandler) CreateUser(c *gin.Context) {
	// bind
	req := &validator.UserCreateRequest{}
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
	}

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err := sh.userController.Create(ctx, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, "success")
}
