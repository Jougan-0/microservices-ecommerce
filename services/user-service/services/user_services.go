package services

import (
	"errors"
	"user-service/models"
	"user-service/repository"
	"user-service/utils"
)

func RegisterUser(user models.User) error {
	hashedPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashedPassword
	return repository.CreateUser(&user)
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
	return token, err
}
