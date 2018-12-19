package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"os"
)

func main() {

	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWARD")
	host := os.Getenv("DATABASE_HOST_NAME")
	dbName := "chatbot"

	dropIfExistsAndCreateDatabase(host, user, password, dbName)
}

func dropIfExistsAndCreateDatabase(
	hostName string,
	userName string,
	password string,
	dbName string) {

	connect := fmt.Sprintf("%s:%s@tcp(%s)/", userName, password, hostName)
	db, _ := gorm.Open("mysql", connect)

	if existsDatabase(*db, dbName) {
		dropDatabase(*db, dbName)
	}

	createDatabase(*db, dbName)
}

func existsDatabase(db gorm.DB, dbName string) bool {

	count := 0
	db.Table("information_schema.schemata").Where("schema_name = ?", dbName).Count(&count)
	exist := count > 0
	return exist
}

func createDatabase(db gorm.DB, dbName string) {
	db.Exec("CREATE DATABASE " + dbName)
}

func dropDatabase(db gorm.DB, dbName string) {
	db.Exec("DROP DATABASE " + dbName)
}
