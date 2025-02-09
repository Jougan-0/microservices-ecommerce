package services

import (
	"errors"
	"microservices/models"
	"microservices/repository"

	"github.com/gofrs/uuid"
)

func AddItemToCart(cart models.Cart) error {
	product, err := repository.GetProductByID((cart.ProductID))
	if err != nil {
		return errors.New("product not found")
	}
	if cart.Quantity > product.Stock {
		return errors.New("insufficient stock")
	}
	return repository.AddToCart(&cart)
}

func GetCart(userID uuid.UUID) ([]models.CartResponse, error) {
	cartItems, err := repository.GetUserCart(userID)
	if err != nil {
		return nil, err
	}
	var validCart []models.CartResponse
	for _, item := range cartItems {
		product, err := repository.GetProductByID((item.ProductID))
		if err == nil && product.Stock > 0 {
			validCart = append(validCart, models.CartResponse{
				ID:          item.ID,
				ProductName: product.Name,
				ProductID:   product.ID,
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

func DeleteCartItem(cartID uuid.UUID, quantity int, userId uuid.UUID) error {
	return repository.RemoveCartItem(cartID, quantity, userId)
}

func ClearUserCart(userID uuid.UUID) error {
	return repository.ClearCart(userID)
}
