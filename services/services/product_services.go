package services

import (
	"microservices/models"
	"microservices/repository"
)

func AddProduct(product models.Product) error {
	return repository.CreateProduct(&product)
}

func ListProducts() ([]models.Product, error) {
	return repository.GetAllProducts()
}

func FindProduct(id uint) (models.Product, error) {
	return repository.GetProductByID(id)
}

func ModifyProduct(product models.Product) error {
	return repository.UpdateProduct(&product)
}

func RemoveProduct(id uint) error {
	return repository.DeleteProduct(id)
}
