package database

import "AltTube-Go/models"

func AddVideo(video models.Video) error {
	dbInstance.Create(&video)
	return nil
}
