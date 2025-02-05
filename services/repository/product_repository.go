package repository

import (
	"microservices/db"
	"microservices/models"
)

func CreateProduct(product *models.Product) error {
	return db.DB.Create(product).Error
}

func GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	err := db.DB.Find(&products).Error
	return products, err
}

func GetProductByID(id uint) (models.Product, error) {
	var product models.Product
	err := db.DB.First(&product, id).Error
	return product, err
}

func UpdateProduct(product *models.Product) error {
	return db.DB.Save(product).Error
}

func DeleteProduct(id uint) error {
	return db.DB.Delete(&models.Product{}, id).Error
}
