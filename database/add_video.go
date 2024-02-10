package database

import "AltTube-Go/model"

func AddVideo(video model.Video) {
	dbInstance.Create(&video)
}
