package main

import (
	"chatbot-back/app"
	"fmt"
	_ "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	_ "net/http"
	"os"
)

func main() {

	err := godotenv.Load(fmt.Sprintf("envfiles/%s.env", "develop"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	requesturl := os.Getenv("RECRUIT_SMARTTALK_ENDPOINT")
	log.Println(requesturl)
	apikey := os.Getenv("RECRUIT_SMARTTALK_API_KEY")
	log.Println(apikey)
	credential := os.Getenv("FIREBASE_ADMIN_SDK_FILENAME")
	log.Println(credential)

	app.Start()
}
