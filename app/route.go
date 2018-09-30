package app

import (
	"github.com/gin-gonic/gin"
	"github.com/kannan-xiao4/chatbot-back/handlers"
)

func RouteV1(app *gin.Engine) {
	helloHandler := handlers.NewHelloHandler()
	welcomeHandler := handlers.NewWelcomeHandler()
	apiGroup := app.Group("api")
	{
		apiGroup.GET("/user/:name", helloHandler.GetName)
		apiGroup.GET("/welcome", welcomeHandler.GetName)
	}

	registerHandler := handlers.NewRegisterHandler()
	userGroup := app.Group("user")
	{
		userGroup.POST("/register", registerHandler.RegisterUser)
		userGroup.POST("/login", registerHandler.LoginUser)
	}

}