package middleware

import (
	"microservices/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func JWTAuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		utils.Logger.Warn("Unauthorized access: Missing Authorization header")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	tokenString := strings.Split(authHeader, " ")
	if len(tokenString) != 2 {
		utils.Logger.Warn("Unauthorized access: Invalid token format")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token format"})
	}
	token, err := jwt.Parse(tokenString[1], func(token *jwt.Token) (interface{}, error) {
		return []byte(utils.GetEnv("JWT_SECRET", "defaultsecret")), nil
	})
	if err != nil || !token.Valid {
		utils.Logger.Warn("Unauthorized access: Invalid token")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		utils.Logger.Warn("❌ Unauthorized: Invalid token claims")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token claims"})
	}

	email, emailOk := claims["email"].(string)
	if !emailOk {
		utils.Logger.Warn("❌ Unauthorized: Email claim missing in token")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token claims"})
	}
	c.Locals("email", email)
	utils.Logger.Info("User authorized successfully")
	return c.Next()
}
