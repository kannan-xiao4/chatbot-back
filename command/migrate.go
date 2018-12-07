package main

import (
	"chatbot-back/db"
	"chatbot-back/models"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := database.Connect()
	db.CreateTable(&models.User{})
	db.CreateTable(&models.ChatMessage{})
}
