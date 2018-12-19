package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"os"
)

func Connect() *gorm.DB {
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWARD")
	host := os.Getenv("DATABASE_HOST_NAME")
	dbName := "chatbot"
	connect := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True", user, password, host, dbName)

	db, err := gorm.Open("mysql", connect)

	if err != nil {
		panic(fmt.Sprintf("fail to connect %s database, connection format is %s, ", dbName, connect) + err.Error())
	}

	return db
}
