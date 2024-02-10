package database

import "AltTube-Go/model"

func GetVideoByV(id string) (*model.Video, error) {
	result := model.Video{}
	dbResult := dbInstance.Where("id = ?", id).First(&result)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return &result, nil
}
