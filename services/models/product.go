package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Product struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name        string    `gorm:"size:100;unique;not null" json:"name"`
	Description string    `gorm:"not null" json:"description"`
	Price       float64   `gorm:"not null" json:"price"`
	Stock       int       `gorm:"not null" json:"stock"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}
