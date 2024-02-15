package models

import (
	"gorm.io/gorm"
	"time"
)

type RefreshToken struct {
	gorm.Model
	Token  string `gorm:"index:idx_token,unique"`
	UserID string `gorm:"index"`
	Expiry time.Time
}
