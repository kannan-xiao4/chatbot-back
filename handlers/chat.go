package handlers

import (
	"chatbot-back/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type chatHandlerInterface interface {
	GetMessages(c *gin.Context)
	AddMessage(c *gin.Context)
}

func NewChatHandler() chatHandlerInterface {
	return &chatHandler{}
}

type chatHandler struct {
}

func (h *chatHandler) GetMessages(c *gin.Context) {
	n := c.Param("user_id")

	id, err := strconv.ParseInt(n, 10, 64)
	if err != nil {
		c.JSON(400, err)
		return
	}

	messages := repository.NewChatMessageRepository().GetByUserID(id)

	if messages == nil {
		c.JSON(404, gin.H{})
		return
	}

	c.JSON(200, messages)
}

func (h *chatHandler) AddMessage(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)
}
