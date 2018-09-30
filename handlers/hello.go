package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type helloHandlerInterface interface {
	GetName(c *gin.Context)
}

func NewHelloHandler() helloHandlerInterface {
	return &helloHandler{}
}

type helloHandler struct {
}

func (h *helloHandler) GetName(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)
}