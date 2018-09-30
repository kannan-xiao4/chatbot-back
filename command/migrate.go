package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/kannan-xiao4/chatbot-back/db"
	"github.com/kannan-xiao4/chatbot-back/models"
)

func main(){
	db := database.Connect();
	db.CreateTable(&models.User{})
}