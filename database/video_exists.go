package database

import "AltTube-Go/models"

func VideoExists(id string) bool {
	result := models.Video{}
	dbResult := dbInstance.Where("id = ?", id).First(&result)
	if dbResult.Error != nil {
		return false
	}
	return true
}
