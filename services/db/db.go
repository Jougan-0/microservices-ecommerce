package db

import (
	"fmt"
	"microservices/models"
	"microservices/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		utils.GetEnv("DB_HOST", "host"),
		utils.GetEnv("DB_USER", "user"),
		utils.GetEnv("DB_PASSWORD", "admin"),
		utils.GetEnv("DB_NAME", "pgsql"),
		utils.GetEnv("DB_PORT", "5431"),
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		utils.Logger.Fatalf("❌ Failed to connect to database: %v", err)
	}

	utils.Logger.Info("✅ Connected to PostgreSQL successfully!")

	DB.AutoMigrate(&models.User{})
}
