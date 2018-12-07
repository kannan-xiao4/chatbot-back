package main

import (
	"chatbot-back/app"
	_ "github.com/gin-gonic/gin"
	_ "net/http"
)

func main() {
	app.Start()
}
