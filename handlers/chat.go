package handlers

import (
	"chatbot-back/repository"
	"github.com/gin-gonic/gin"
	"net/http"
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

	user := repository.NewUserRepository().GetByUID(c.GetString("uid"))

	if user == nil {
		c.JSON(404, gin.H{"error": http.StatusText(http.StatusNotFound)})
		return
	}

	messages := repository.NewChatMessageRepository().GetByUserID(user.Id)

	if messages == nil {
		c.JSON(404, gin.H{"error": http.StatusText(http.StatusNotFound)})
		return
	}

	c.JSON(200, messages)
}

func (h *chatHandler) AddMessage(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)
}
