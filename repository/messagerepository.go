package repository

import (
	"chatbot-back/db"
	"chatbot-back/request"
	"golang.org/x/net/context"
	"log"
)

type MessageRepository struct {
}

func NewMessageRepository() MessageRepository {
	return MessageRepository{}
}

func (m MessageRepository) PushMessage(message request.PostMessage) {
	db := database.FirebaseConnect("message")
	_, err := db.Push(context.Background(), message)

	if err != nil {
		log.Fatalln("Error pushing child node:", err)
	}
}
