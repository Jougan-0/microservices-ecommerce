package repository

import (
	"context"
	"user-service/db"
	"user-service/models"
)

func CreateUser(user models.User) error {
	_, err := db.DB.Exec(context.Background(),
		"INSERT INTO users (name, email, password_hash) VALUES ($1, $2, $3)",
		user.Name, user.Email, user.Password)
	return err
}

func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := db.DB.QueryRow(context.Background(),
		"SELECT id, name, email, password_hash, created_at FROM users WHERE email=$1", email).
		Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return user, err
}
