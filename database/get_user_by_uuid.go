package database

import "AltTube-Go/models"

func GetUserByID(id string) (*models.User, error) {
	result := models.User{}
	dbResult := dbInstance.Where("id = ?", id).First(&result)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return &result, nil
}
