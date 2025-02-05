package handlers

import (
	"user-service/models"
	"user-service/services"
	"user-service/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func RegisterUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		utils.Logger.WithError(err).Warn("Invalid request payload")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := services.RegisterUser(user); err != nil {
		utils.Logger.WithError(err).Error("User registration failed")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "User registration failed"})
	}

	utils.Logger.WithFields(logrus.Fields{
		"email": user.Email,
	}).Info("User registered successfully via API")
	return c.JSON(fiber.Map{"message": "User registered successfully"})
}

func LoginUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		utils.Logger.WithError(err).Warn("Invalid login request")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	token, err := services.LoginUser(user.Email, user.Password)
	if err != nil {
		utils.Logger.WithFields(logrus.Fields{
			"email": user.Email,
		}).Warn("User login attempt failed")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	utils.Logger.WithFields(logrus.Fields{
		"email": user.Email,
	}).Info("User logged in successfully via API")
	return c.JSON(fiber.Map{"token": token})
}
