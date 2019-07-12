package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
