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

// GetByUID ...
func (m UserRepository) GetByUID(uid string) *models.User {
	var user = models.User{Uuid: uid}

	db := database.Connect()
	if db.First(&user).RecordNotFound() {
		return nil
	}
	return &user
}

func (m UserRepository) Persist(user *models.User) *models.User {

	db := database.Connect()

	if db.First(&user).RecordNotFound() {
		db.Create(&user)
	} else {
		db.Update(&user)
	}
	return user
}
