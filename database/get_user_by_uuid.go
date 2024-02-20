package database

import "AltTube-Go/models"

func GetUserByUUID(uuid string) (*models.User, error) {
	result := models.User{}
	dbResult := dbInstance.Where("uuid = ?", uuid).First(&result)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return &result, nil
}
