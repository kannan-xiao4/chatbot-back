package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type welcomeHandlerInterface interface {
	GetName(c *gin.Context)
}

func NewWelcomeHandler() welcomeHandlerInterface {
	return &welcomeHandler{}
}

type welcomeHandler struct {
}

func (h *welcomeHandler) GetName(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")

	uid := c.GetString("uid")

	c.String(http.StatusOK, "Hello %s %s", firstname, uid)
}
