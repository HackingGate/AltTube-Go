package database

import (
	"AltTube-Go/models"
	"errors"

	"gorm.io/gorm"
)

func AddLikeVideo(user *models.User, video *models.Video) error {
	like := models.LikeVideo{
		UserID:  user.ID,
		User:    *user,
		VideoID: video.ID,
		Video:   *video,
	}
	return dbInstance.Create(&like).Error
}

func ReadIsLikedVideo(user *models.User, video *models.Video) (bool, error) {
	var like models.LikeVideo
	dbResult := dbInstance.Where("user_id = ? AND video_id = ?", user.ID, video.ID).First(&like)
	if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if dbResult.Error != nil {
		return false, dbResult.Error
	}
	return true, nil
}

func RemoveLikeVideo(user *models.User, video *models.Video) error {
	return dbInstance.Where("user_id = ? AND video_id = ?", user.ID, video.ID).Delete(&models.LikeVideo{}).Error
}

func RemoveAllLikesByUserID(userID string) error {
	return dbInstance.Unscoped().Where("user_id = ?", userID).Delete(&models.LikeVideo{}).Error
}

func GetAllLikesByUserID(userID string) ([]models.LikeVideo, error) {
	var likes []models.LikeVideo
	dbResult := dbInstance.Where("user_id = ?", userID).Find(&likes)
	return likes, dbResult.Error
}
