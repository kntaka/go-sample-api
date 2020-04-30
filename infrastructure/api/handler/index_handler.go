package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type indexHandler struct {
}

type IndexHandler interface {
	Index(c *gin.Context)
}

func NewIndexHandler() IndexHandler {
	return &indexHandler{}
}

func (ih *indexHandler) Index(c *gin.Context) {
	c.JSON(http.StatusOK, "Success")
}
