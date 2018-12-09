package app

import (
	"chatbot-back/handlers"
	"github.com/gin-gonic/gin"
)

func RouteV1(app *gin.Engine) {
	helloHandler := handlers.NewHelloHandler()
	welcomeHandler := handlers.NewWelcomeHandler()
	chatHandler := handlers.NewChatHandler()
	apiGroup := app.Group("api")
	{
		apiGroup.GET("/user/:name", helloHandler.GetName)
		apiGroup.GET("/welcome", welcomeHandler.GetName)
		apiGroup.GET("/chat/:user_id", chatHandler.GetMessages)
	}

	registerHandler := handlers.NewRegisterHandler()
	userGroup := app.Group("user")
	{
		userGroup.POST("/register", registerHandler.RegisterUser)
		userGroup.POST("/login", registerHandler.LoginUser)
	}

}
