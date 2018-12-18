package handlers

import (
	"chatbot-back/models"
	"chatbot-back/repository"
	"chatbot-back/request"
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
		c.JSON(http.StatusNotFound, gin.H{"error": http.StatusText(http.StatusNotFound)})
		return
	}

	messages := repository.NewChatMessageRepository().GetByUserID(user.Id)

	if messages == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": http.StatusText(http.StatusNotFound)})
		return
	}

	c.JSON(http.StatusOK, messages)
}

func (h *chatHandler) AddMessage(c *gin.Context) {

	postMessage := request.PostMessage{}

	err := c.BindJSON(&postMessage)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}

	userId := c.GetString("uid")
	user := repository.NewUserRepository().GetByUID(userId)

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": http.StatusText(http.StatusNotFound)})
		return
	}

	meg := repository.NewMessageRepository()
	meg.PushMessage(postMessage)

	message := models.ChatMessage{UserId: userId, Message: postMessage.Message}

	message.Response = "まだ未実装ですの"

	repository.NewChatMessageRepository().Persist(message)

	meg.PushMessage(request.PostMessage{
		Message:  message.Response,
		UserName: "bot",
		Image:    "none",
	})

	c.JSON(http.StatusOK, message)
}
