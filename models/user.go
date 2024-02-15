package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID   string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string
}

type EditEmail struct {
	Email string `json:"email"`
}
