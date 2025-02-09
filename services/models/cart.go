package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Cart struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	ProductID uuid.UUID `gorm:"type:uuid;not null" json:"product_id"`
	Quantity  int       `gorm:"not null" json:"quantity"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
type CartResponse struct {
	ID          uuid.UUID `json:"id"`
	ProductID   uuid.UUID `json:"product_id"`
	ProductName string    `json:"product_name"`
	Quantity    int       `json:"quantity"`
	TotalPrice  float64   `json:"total_price"`
}
