package database

import "AltTube-Go/models"

func GetRefreshTokenByUserID(userID string) ([]models.RefreshToken, error) {
	var result []models.RefreshToken
	dbResult := dbInstance.Where("user_id = ?", userID).Find(&result)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return result, nil
}
