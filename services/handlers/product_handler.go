package handlers

import (
	"microservices/models"
	"microservices/services"

	"github.com/gofiber/fiber/v2"
)

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := services.AddProduct(product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Product creation failed"})
	}

	return c.JSON(fiber.Map{"message": "Product added successfully"})
}

func GetProducts(c *fiber.Ctx) error {
	products, err := services.ListProducts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch products"})
	}

	return c.JSON(products)
}
func UpdateProduct(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	product, err := services.FindProduct(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
	}
	var updateData models.Product
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	if updateData.Name != "" {
		product.Name = updateData.Name
	}
	if updateData.Description != "" {
		product.Description = updateData.Description
	}
	if updateData.Price != 0 {
		product.Price = updateData.Price
	}
	if updateData.Stock != 0 {
		product.Stock = updateData.Stock
	}
	if err := services.ModifyProduct(product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Product update failed"})
	}

	return c.JSON(fiber.Map{"message": "Product updated successfully"})

}

func DeleteProduct(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	err := services.RemoveProduct(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Product deletion failed"})
	}

	return c.JSON(fiber.Map{"message": "Product deleted successfully"})
}
func GetProductByID(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	product, err := services.FindProduct(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
	}

	return c.JSON(product)
}
