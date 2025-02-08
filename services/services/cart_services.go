package services

import (
	"errors"
	"microservices/models"
	"microservices/repository"
)

func AddItemToCart(cart models.Cart) error {
	product, err := repository.GetProductByID(cart.ProductID)
	if err != nil {
		return errors.New("product not found")
	}
	if cart.Quantity > product.Stock {
		return errors.New("insufficient stock")
	}
	return repository.AddToCart(&cart)
}

func GetCart(userID uint) ([]models.CartResponse, error) {
	cartItems, err := repository.GetUserCart(userID)
	if err != nil {
		return nil, err
	}
	var validCart []models.CartResponse
	for _, item := range cartItems {
		product, err := repository.GetProductByID(item.ProductID)
		if err == nil && product.Stock > 0 {
			validCart = append(validCart, models.CartResponse{
				ID:          item.ID,
				ProductName: product.Name,
				Quantity:    item.Quantity,
				TotalPrice:  product.Price * float64(item.Quantity),
			})
		}
	}

	return validCart, nil
}

func ModifyCartItem(cart models.Cart) error {
	return repository.UpdateCartItem(&cart)
}

func DeleteCartItem(cartID uint, quantity int, userId uint) error {
	return repository.RemoveCartItem(cartID, quantity, userId)
}

func ClearUserCart(userID uint) error {
	return repository.ClearCart(userID)
}
