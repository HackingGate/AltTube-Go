package models

import "gorm.io/gorm"

type LikeVideo struct {
	gorm.Model
	UserID  string
	User    User   `gorm:"foreignKey:UserID"`
	VideoID string `gorm:"index"`
	Video   Video  `gorm:"foreignKey:VideoID"`
}

type LikeVideoRequest struct {
	VideoID string `json:"video_id" binding:"required"`
}

type LikeVideoResponse struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	ThumbnailUrl string `json:"thumbnail_url"`
}
