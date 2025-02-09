package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if os.Getenv("BUILD_TYPE") == "dockerfile" {
		fmt.Println("🔸 Detected Docker build environment; skipping .env load.")
		return
	}
	err := godotenv.Load()
	if err != nil {
		fmt.Println("❌ Error loading .env file. Using system environment variables.")
	} else {
		fmt.Println("✅ .env file loaded successfully")
	}
}
func GetEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
