package repository

import (
	"errors"
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

func RemoveCartItem(cartID uint, quantity int, userId uint) error {
	var cart models.Cart
	err := db.DB.Where("id = ?", cartID).First(&cart).Error
	if err != nil {
		return err
	}
	if cart.UserID != userId {
		return errors.New("cart item not found")
	}

	if quantity >= cart.Quantity {
		return db.DB.Delete(&models.Cart{}, cartID).Error
	}

	return db.DB.Model(&cart).Update("quantity", cart.Quantity-quantity).Error
}

func ClearCart(userID uint) error {
	return db.DB.Where("user_id = ?", userID).Delete(&models.Cart{}).Error
}
