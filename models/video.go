package models

import (
	"gorm.io/gorm"
	"time"
)

type Video struct {
	ID           string    `gorm:"primaryKey;unique"`
	Title        string    `gorm:"not null"`
	Description  string    `gorm:"not null"`
	UploadDate   time.Time `gorm:"not null"`
	Uploader     string    `gorm:"not null"`
	UploaderUrl  string    `gorm:"not null"`
	ThumbnailUrl string    `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
