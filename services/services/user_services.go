package services

import (
	"errors"
	"microservices/models"
	"microservices/repository"
	"microservices/utils"

	"github.com/sirupsen/logrus"
)

func RegisterUser(user models.User) error {
	hashedPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashedPassword
	err := repository.CreateUser(&user)
	if err != nil {
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
		return "", errors.New("user not found")
	}

	if !utils.CheckPasswordHash(password, storedUser.Password) {
		return "", errors.New("invalid password")
	}

	token, err := utils.GenerateJWT(storedUser.Email)
	if err != nil {
		return "", err
	}

	utils.Logger.WithFields(logrus.Fields{
		"email": email,
	}).Info("User logged in successfully")
	return token, nil
}
func FetchUserProfile(email string) (*models.User, error) {
	user, err := repository.GetUserByEmail(email)
	if err != nil {
		utils.Logger.WithFields(logrus.Fields{"email": email}).Warn("User profile not found")
		return nil, err
	}
	return user, nil
}

func UpdateUserProfile(email string, updateData models.UserUpdate) error {
	user, err := repository.GetUserByEmail(email)
	if err != nil {
		return errors.New("user not found")
	}

	user.Name = updateData.Name
	err = repository.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(email string) error {
	err := repository.DeleteUser(email)
	if err != nil {
		return err
	}
	return nil
}
