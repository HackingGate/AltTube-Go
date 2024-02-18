package database

import "AltTube-Go/models"

func RemoveRefreshTokensByID(ids []uint) error {
	return dbInstance.Where("id IN ?", ids).Delete(&models.RefreshToken{}).Error
}
