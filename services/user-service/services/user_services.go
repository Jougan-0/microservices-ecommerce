package services

import (
	"user-service/models"
	"user-service/repository"
	"user-service/utils"
)

func RegisterUser(user models.User) error {
	hashedPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashedPassword
	return repository.CreateUser(user)
}

func LoginUser(email, password string) (string, error) {
	storedUser, err := repository.GetUserByEmail(email)
	if err != nil || !utils.CheckPasswordHash(password, storedUser.Password) {
		return "", err
	}

	token, err := utils.GenerateJWT(storedUser.Email)
	return token, err
}
