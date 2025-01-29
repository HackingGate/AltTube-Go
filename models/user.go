package models

type User struct {
	Entity
	ID       string `gorm:"primaryKey;unique"`
	Email    string `gorm:"unique"`
	Password string `gorm:"not null"`
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
