package models

import (
	_ "regexp"
	"time"
)

type User struct {
	Id int64
	Name string `sql:"size:255" form:"name" json:"name" binding:"required,name"`
	Uuid string `sql:"size:32" form:"address" json:"uuid" binding:"required,uuid"`
	Password string `sql:"size:255" form:"password" json:"password" binding:"required,password"`
	CreatedAt time.Time
	UpdatedAT time.Time
}