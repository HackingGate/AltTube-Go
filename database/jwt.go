package database

import (
	"AltTube-Go/models"
	"time"
)

// AddRefreshToken creates and stores a new refresh token in the database.
func AddRefreshToken(token, userID string, expiry time.Time) error {
	refreshToken := models.RefreshToken{
		Token:  token,
		UserID: userID,
		Expiry: expiry,
	}
	return dbInstance.Create(&refreshToken).Error
}

// ValidateRefreshToken checks if the given token exists and is not expired.
func ValidateRefreshToken(token string) (string, bool) {
	var refreshToken models.RefreshToken
	result := dbInstance.Where("token = ? AND expiry > ?", token, time.Now()).First(&refreshToken)
	if result.Error != nil || result.RowsAffected == 0 {
		return "", false // Token not found or expired
	}
	return refreshToken.UserID, true
}

// RemoveRefreshToken deletes a refresh token from the database.
func RemoveRefreshToken(token string) error {
	result := dbInstance.Unscoped().Where("token = ?", token).Delete(&models.RefreshToken{})
	return result.Error
}

func GetRefreshTokenByUserID(userID string) (string, error) {
	var refreshToken models.RefreshToken
	result := dbInstance.Where("user_id = ? AND expiry > ?", userID, time.Now()).First(&refreshToken)
	return refreshToken.Token, result.Error
}