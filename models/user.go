package models

import "time"

type User struct {
	ID        string `gorm:"primaryKey;unique"`
	Email     string `gorm:"unique"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index"`
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
