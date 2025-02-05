package services

import (
	"errors"
	"fmt"
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

func GetCart(userID uint) ([]models.Cart, error) {
	cartItems, err := repository.GetUserCart(userID)
	if err != nil {
		return nil, err
	}
	var validCart []models.Cart
	for _, item := range cartItems {
		fmt.Println("item.ProductID", item.ProductID)
		product, err := repository.GetProductByID(item.ProductID)
		if err == nil && product.Stock > 0 {
			validCart = append(validCart, item)
		}
	}

	return validCart, nil
}

func ModifyCartItem(cart models.Cart) error {
	return repository.UpdateCartItem(&cart)
}

func DeleteCartItem(cartID uint) error {
	return repository.RemoveCartItem(cartID)
}

func ClearUserCart(userID uint) error {
	return repository.ClearCart(userID)
}
