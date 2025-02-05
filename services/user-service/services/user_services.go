package services

import (
	"errors"
	"user-service/models"
	"user-service/repository"
	"user-service/utils"

	"github.com/sirupsen/logrus"
)

func RegisterUser(user models.User) error {
	hashedPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashedPassword
	err := repository.CreateUser(&user)
	if err != nil {
		utils.Logger.WithError(err).Error("User registration failed")
		return err
	}
	utils.Logger.WithFields(logrus.Fields{
		"email": user.Email,
	}).Info("User registered successfully")
	return nil
}

func LoginUser(email, password string) (string, error) {
	storedUser, err := repository.GetUserByEmail(email)
	if err != nil {
		utils.Logger.WithFields(logrus.Fields{
			"email": email,
		}).Warn("User login failed: User not found")
		return "", errors.New("user not found")
	}

	if !utils.CheckPasswordHash(password, storedUser.Password) {
		utils.Logger.WithFields(logrus.Fields{
			"email": email,
		}).Warn("User login failed: Invalid password")
		return "", errors.New("invalid password")
	}

	token, err := utils.GenerateJWT(storedUser.Email)
	if err != nil {
		utils.Logger.WithError(err).Error("Failed to generate JWT")
		return "", err
	}

	utils.Logger.WithFields(logrus.Fields{
		"email": email,
	}).Info("User logged in successfully")
	return token, nil
}
