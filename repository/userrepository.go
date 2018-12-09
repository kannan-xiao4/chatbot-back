package repository

import (
	"chatbot-back/db"
	"chatbot-back/models"
)

// UserRepository is
type UserRepository struct {
}

// NewUserRepository ...
func NewUserRepository() UserRepository {
	return UserRepository{}
}

// GetByID ...
func (m UserRepository) GetByID(id int64) *models.User {
	var user = models.User{Id: id}

	db := database.Connect()
	if db.First(&user).RecordNotFound() {
		return nil
	}
	return &user
}
