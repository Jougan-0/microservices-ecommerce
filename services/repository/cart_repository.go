package repository

import (
	"errors"
	"microservices/db"
	"microservices/models"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

func AddToCart(cart *models.Cart) error {
	var existingCart models.Cart
	err := db.DB.Where("user_id = ? AND product_id = ?", cart.UserID, cart.ProductID).First(&existingCart).Error

	if err == nil {
		existingCart.Quantity += cart.Quantity
		return db.DB.Save(&existingCart).Error
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return db.DB.Create(cart).Error
	}
	return err
}

func GetUserCart(userID uuid.UUID) ([]models.Cart, error) {
	var cart []models.Cart
	err := db.DB.Where("user_id = ?", userID).Find(&cart).Error
	return cart, err
}

func UpdateCartItem(cart *models.Cart) error {
	return db.DB.Save(cart).Error
}

func RemoveCartItem(cartID uuid.UUID, quantity int, userId uuid.UUID) error {
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

func ClearCart(userID uuid.UUID) error {
	return db.DB.Where("user_id = ?", userID).Delete(&models.Cart{}).Error
}
