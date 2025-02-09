package services

import (
	"microservices/models"
	"microservices/repository"

	"github.com/gofrs/uuid"
)

func AddProduct(product models.Product) error {
	return repository.CreateProduct(&product)
}

func ListProducts() ([]models.Product, error) {
	return repository.GetAllProducts()
}

func FindProduct(id uuid.UUID) (models.Product, error) {
	return repository.GetProductByID(id)
}

func ModifyProduct(product models.Product) error {
	return repository.UpdateProduct(&product)
}

func RemoveProduct(id uuid.UUID) error {
	return repository.DeleteProduct(id)
}
