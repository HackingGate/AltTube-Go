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

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Signup struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type EditEmail struct {
	Email string `json:"email"`
}
