package database

import "AltTube-Go/model"

func VideoExists(id string) bool {
	result := model.Video{}
	dbResult := dbInstance.Where("id = ?", id).First(&result)
	if dbResult.Error != nil {
		return false
	}
	return true
}
