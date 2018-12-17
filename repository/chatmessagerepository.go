package repository

import (
	"chatbot-back/db"
	"chatbot-back/models"
)

type ChatMessageRepository struct {
}

// NewUserRepository ...
func NewChatMessageRepository() ChatMessageRepository {
	return ChatMessageRepository{}
}

// GetByID ...
func (m ChatMessageRepository) GetByID(id int64) *models.ChatMessage {
	var message = models.ChatMessage{Id: id}

	db := database.Connect()
	if db.First(&message).RecordNotFound() {
		return nil
	}
	return &message
}

func (m ChatMessageRepository) GetByUserID(id int64) *[]models.ChatMessage {
	var messages []models.ChatMessage

	db := database.Connect()
	if db.Find(&messages, "user_id=?", id).RecordNotFound() {
		return nil
	}
	return &messages
}

func (m ChatMessageRepository) Persist(message models.ChatMessage) *models.ChatMessage {

	db := database.Connect()

	if db.Create(&message).RecordNotFound() {
		return nil
	}
	return &message
}
