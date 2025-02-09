package handlers

import (
	"microservices/models"
	"microservices/repository"
	"microservices/services"
	"microservices/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

func AddToCart(c *fiber.Ctx) error {
	var cart models.Cart
	if err := c.BodyParser(&cart); err != nil {
		utils.Logger.WithError(err).Warn("Invalid request payload")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	email := c.Locals("email").(string)
	user, err := repository.GetUserByEmail(email)
	if err != nil {
		utils.Logger.WithError(err).Error("Failed to fetch user")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch user"})
	}
	cart.UserID = user.ID
	if err := services.AddItemToCart(cart); err != nil {
		utils.Logger.WithError(err).Error("Failed to add item to cart")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to add item to cart"})
	}

	utils.Logger.WithFields(logrus.Fields{
		"user_id":    cart.UserID,
		"product_id": cart.ProductID,
		"quantity":   cart.Quantity,
	}).Info("Item added to cart")

	return c.JSON(fiber.Map{"message": "Item added to cart"})
}

func GetCart(c *fiber.Ctx) error {
	email := c.Locals("email").(string)
	user, err := repository.GetUserByEmail(email)
	if err != nil {
		utils.Logger.WithError(err).Error("Failed to fetch user")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch user"})
	}
	cart, err := services.GetCart(user.ID)
	if err != nil {
		utils.Logger.WithError(err).Error("Failed to fetch cart")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch cart"})
	}

	utils.Logger.WithField("user_id", user.ID).Info("Fetched cart successfully")
	if len(cart) == 0 {
		return c.JSON(fiber.Map{"message": "Cart is empty"})
	}
	return c.JSON(cart)
}

func RemoveCartItem(c *fiber.Ctx) error {
	quantity, _ := strconv.Atoi(c.Query("quantity"))
	email := c.Locals("email").(string)
	//extra security
	user, err := repository.GetUserByEmail(email)
	if err != nil {
		utils.Logger.WithError(err).Error("Failed to fetch user")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch user"})
	}
	cartId, _ := uuid.FromString(c.Params("id"))
	err = services.DeleteCartItem(cartId, int(quantity), user.ID)
	if err != nil {
		utils.Logger.WithError(err).Error("Failed to remove item from cart")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to remove item"})
	}

	utils.Logger.WithField("cart_id", cartId).Info("Item removed from cart")
	return c.JSON(fiber.Map{"message": "Item removed from cart"})
}

func ClearCart(c *fiber.Ctx) error {
	email := c.Locals("email").(string)
	user, err := repository.GetUserByEmail(email)
	if err != nil {
		utils.Logger.WithError(err).Error("Failed to fetch user")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch user"})
	}
	err = services.ClearUserCart(user.ID)
	if err != nil {
		utils.Logger.WithError(err).Error("Failed to clear cart")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to clear cart"})
	}

	utils.Logger.WithField("user_id", user.ID).Info("Cart cleared successfully")
	return c.JSON(fiber.Map{"message": "Cart cleared"})
}
