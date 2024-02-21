package models

import (
	"gorm.io/gorm"
	"time"
)

type AccessToken struct {
	gorm.Model
	Token          string `gorm:"index:idx_token"`
	UserID         string `gorm:"index"`
	User           User   `gorm:"foreignKey:UserID"`
	Expiry         time.Time
	RefreshTokenID uint         `gorm:"index"`
	RefreshToken   RefreshToken `gorm:"foreignKey:RefreshTokenID"`
}

type RefreshToken struct {
	gorm.Model
	Token     string `gorm:"index:idx_token"`
	Expiry    time.Time
	UserAgent string
	IPAddress string
	UserID    string `gorm:"index"`
	User      User   `gorm:"foreignKey:UserID"`
}
