package database

import (
	"firebase.google.com/go"
	"firebase.google.com/go/db"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

func FirebaseConnect(path string) *db.Ref {

	ctx := context.Background()
	conf := &firebase.Config{DatabaseURL: "https://chatbot-base.firebaseio.com/"}

	// Firebase SDK のセットアップ
	opt := option.WithCredentialsFile("credentials/chatbot-base-firebase-adminsdk-b3c63-219633f486.json")
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		panic(fmt.Sprintf("error initializing app: %v", err))

		return nil
	}

	cliant, err := app.Database(ctx)

	if err != nil {
		panic(fmt.Sprintf("Error initializing database client: %v", err))
		return nil
	}

	return cliant.NewRef(path)
}
