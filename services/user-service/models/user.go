package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:50;not null" json:"name"`
	Email     string    `gorm:"size:100;unique;not null" json:"email"`
	Password  string    `gorm:"column:password_hash;not null" json:"-"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
