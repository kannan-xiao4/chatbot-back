package middleware

import (
	"context"
	"firebase.google.com/go"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"log"
	"net/http"
)

func FirebaseAuth(c *gin.Context) {

	// Firebase SDK のセットアップ
	opt := option.WithCredentialsFile("credentials/chatbot-base-firebase-adminsdk-b3c63-219633f486.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		fmt.Printf("error initializing app: %v", err)
		firebaseUninitialized(c)
		return
	}
	auth, err := app.Auth(context.Background())
	if err != nil {
		fmt.Printf("error initializing app: %v", err)
		firebaseUninitialized(c)
		return
	}

	authorizationHeader := c.Request.Header.Get("Authorization")
	if authorizationHeader == "" {
		unauthorized(c)
		return
	}
	if authorizationHeader[:7] != "Bearer " {
		unauthorized(c)
		return
	}
	tokenString := authorizationHeader[7:]

	// JWT の検証
	token, err := auth.VerifyIDToken(context.Background(), tokenString)
	if err != nil {
		// JWT が無効なら Handler に進まず別処理
		fmt.Printf("error verifying ID token: %v\n", err)
		unauthorized(c)
		return
	}
	log.Printf("Verified ID token: %v\n", token)
}

func unauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"error": http.StatusText(http.StatusUnauthorized),
	})
	c.Abort()
}

func firebaseUninitialized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"error": http.StatusText(http.StatusInternalServerError),
	})
	c.Abort()
}
