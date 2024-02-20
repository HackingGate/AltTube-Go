package models

import (
	"gorm.io/gorm"
	"time"
)

type RefreshToken struct {
	gorm.Model
	Token     string `gorm:"index:idx_token"`
	Expiry    time.Time
	UserAgent string
	IPAddress string
	UserID    string `gorm:"index"`
	User      User   `gorm:"foreignKey:UserID"`
}
