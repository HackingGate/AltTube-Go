package database

import "AltTube-Go/model"

func AddVideo(video model.Video) error {
	dbInstance.Create(&video)
	return nil
}
