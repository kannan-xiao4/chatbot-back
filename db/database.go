package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func Connect() *gorm.DB{
	user := "root"
	password := "secret"
	host := "127.0.0.1:3306"
	dbName := "chatbot"
	connect := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True", user, password, host, dbName)

	db, err := gorm.Open("mysql", connect)

	if err != nil {
		panic(fmt.Sprintf("fail to connect %s database, connection format is %s, ", dbName, connect) + err.Error())
	}

	return db
}