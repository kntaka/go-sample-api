package router

import (
	"github.com/gin-gonic/gin"
	"github.com/takkee/go-sample-api/infrastructure/api/handler"
)

func NewRouter(r *gin.Engine, handler handler.AppHandler) {

	// Index (Might not be used.Just test.)
	r.GET("/", handler.Index)

	// version 1
	v1 := r.Group("/v1")
	{
		// get a my user info
		v1.GET("/users/:id", handler.Index)
		// create a user
		v1.POST("/users", handler.CreateUser)
		// update user
		v1.PUT("/users/:id", handler.Index)
		// delete user
		v1.DELETE("/users/:id", handler.Index)
	}
}
