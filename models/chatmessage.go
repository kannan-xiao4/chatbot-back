package models

import (
	_ "regexp"
	"time"
)

type ChatMessage struct {
	Id        int64
	UserId    int64
	Message   string `sql:"size:255" form:"message" json:"message" binding:"required,message"`
	Response  string `sql:"size:255" form:"response" json:"response" binding:"required,response"`
	CreatedAt time.Time
	UpdatedAT time.Time
}
