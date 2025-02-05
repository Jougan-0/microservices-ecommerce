package repository

import (
	"microservices/db"
	"microservices/models"
	"microservices/utils"

	"github.com/sirupsen/logrus"
)

func CreateUser(user *models.User) error {
	if err := db.DB.Create(user).Error; err != nil {
		return err
	}
	utils.Logger.WithFields(logrus.Fields{
		"email": user.Email,
	}).Info("User created successfully")
	return nil
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := db.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	utils.Logger.WithFields(logrus.Fields{
		"email": email,
	}).Info("User fetched successfully")
	return &user, nil
}
func UpdateUser(user *models.User) error {
	err := db.DB.Save(user).Error
	return err
}

func DeleteUser(email string) error {
	err := db.DB.Where("email = ?", email).Delete(&models.User{}).Error
	return err
}
