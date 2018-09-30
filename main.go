package main

import (
	_ "github.com/gin-gonic/gin"
	"github.com/kannan-xiao4/chatbot-back/app"
	_ "net/http"
)

func main() {
	app.Start()
}
