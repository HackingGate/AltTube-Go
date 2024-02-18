package models

import (
	"gorm.io/gorm"
	"time"
)

type RefreshToken struct {
	gorm.Model
	Token     string `gorm:"index:idx_token"`
	UserID    string `gorm:"index"`
	Expiry    time.Time
	UserAgent string
	IPAddress string
}
