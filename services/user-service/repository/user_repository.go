package repository

import (
	"user-service/db"
	"user-service/models"
)

func CreateUser(user *models.User) error {
	return db.DB.Create(user).Error
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := db.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}
