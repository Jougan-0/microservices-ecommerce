package handlers

import (
	"microservices/models"
	"microservices/services"
	"microservices/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func RegisterUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		utils.Logger.WithError(err).Error("Invalid request payload")
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
		utils.Logger.WithError(err).Error("Invalid login request")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	token, err := services.LoginUser(user.Email, user.Password)
	if err != nil {
		utils.Logger.WithFields(logrus.Fields{
			"email": user.Email,
		}).Error("User login attempt failed")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	utils.Logger.WithFields(logrus.Fields{
		"email": user.Email,
	}).Info("User logged in successfully via API")
	return c.JSON(fiber.Map{"token": token})
}

func GetUserProfile(c *fiber.Ctx) error {
	emailValue := c.Locals("email")
	if emailValue == nil {
		utils.Logger.Error("❌ No email found in request context")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	email, ok := emailValue.(string)
	if !ok {
		utils.Logger.Error("❌ Failed to cast email to string")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token payload"})
	}

	user, err := services.FetchUserProfile(email)
	if err != nil {
		utils.Logger.WithError(err).Error("Failed to fetch user profile")
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}
	return c.JSON(user)
}

func UpdateUserProfile(c *fiber.Ctx) error {
	email := c.Locals("email").(string)
	var updateData models.UserUpdate
	if err := c.BodyParser(&updateData); err != nil {
		utils.Logger.Error("Invalid request payload")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	err := services.UpdateUserProfile(email, updateData)
	if err != nil {
		utils.Logger.WithError(err).Error("Failed to update profile")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update profile"})
	}
	return c.JSON(fiber.Map{"message": "Profile updated successfully"})
}

func DeleteUserAccount(c *fiber.Ctx) error {
	email := c.Locals("email").(string)
	err := services.DeleteUser(email)
	if err != nil {
		utils.Logger.WithError(err).Error("Failed to delete account")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete account"})
	}
	return c.JSON(fiber.Map{"message": "Account deleted successfully"})
}
