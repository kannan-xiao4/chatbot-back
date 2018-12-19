package app

import (
	"chatbot-back/handlers"
	"chatbot-back/middleware"
	"github.com/gin-gonic/gin"
)

func RouteV1(app *gin.Engine) {
	helloHandler := handlers.NewHelloHandler()
	welcomeHandler := handlers.NewWelcomeHandler()
	chatHandler := handlers.NewChatHandler()

	app.Use(middleware.CORS)

	app.GET("/user/:name", helloHandler.GetName)
	app.GET("/welcome", welcomeHandler.GetName)

	apiGroup := app.Group("/api")
	{
		authorized := apiGroup.Group("/", middleware.FirebaseAuth)
		authorized.GET("/chat", chatHandler.GetMessages)
		authorized.POST("/chat", chatHandler.AddMessage)
	}

	registerHandler := handlers.NewRegisterHandler()
	userGroup := app.Group("/user")
	{
		userGroup.POST("/register", registerHandler.RegisterUser)
		userGroup.POST("/login", registerHandler.LoginUser)
	}

}
