package database

import "AltTube-Go/models"

func GetVideoByV(id string) (*models.Video, error) {
	result := models.Video{}
	dbResult := dbInstance.Where("id = ?", id).First(&result)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return &result, nil
}
