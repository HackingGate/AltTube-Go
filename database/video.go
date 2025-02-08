package database

import "github.com/hackinggate/alttube-go/models"

func AddVideo(video models.Video) error {
	dbInstance.Create(&video)
	return nil
}

func GetVideoByV(id string) (*models.Video, error) {
	result := models.Video{}
	dbResult := dbInstance.Where("id = ?", id).First(&result)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return &result, nil
}

func GetUserByID(id string) (*models.User, error) {
	result := models.User{}
	dbResult := dbInstance.Where("id = ?", id).First(&result)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return &result, nil
}

func VideoExists(id string) bool {
	result := models.Video{}
	dbResult := dbInstance.Where("id = ?", id).First(&result)
	if dbResult.Error != nil {
		return false
	}
	return true
}
