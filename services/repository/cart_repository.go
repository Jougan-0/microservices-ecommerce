package repository

import (
	"microservices/db"
	"microservices/models"
)

func AddToCart(cart *models.Cart) error {
	return db.DB.Create(cart).Error
}

func GetUserCart(userID uint) ([]models.Cart, error) {
	var cart []models.Cart
	err := db.DB.Where("user_id = ?", userID).Find(&cart).Error
	return cart, err
}

func UpdateCartItem(cart *models.Cart) error {
	return db.DB.Save(cart).Error
}

func RemoveCartItem(cartID uint) error {
	return db.DB.Delete(&models.Cart{}, cartID).Error
}

func ClearCart(userID uint) error {
	return db.DB.Where("user_id = ?", userID).Delete(&models.Cart{}).Error
}
