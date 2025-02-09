package database

import (
	"context"

	"github.com/hackinggate/alttube-go/ent"
	"github.com/hackinggate/alttube-go/ent/video"
)

// AddVideo adds a new video to the database.
func AddVideo(ctx context.Context, videoToAdd ent.Video) (*ent.Video, error) {
	videoAdded, err := Client.Video.
		Create().
		SetID(videoToAdd.ID).
		SetTitle(videoToAdd.Title).
		SetDescription(videoToAdd.Description).
		SetUploadDate(videoToAdd.UploadDate).
		SetUploader(videoToAdd.Uploader).
		SetUploaderUrl(videoToAdd.UploaderUrl).
		SetThumbnailUrl(videoToAdd.ThumbnailUrl).
		Save(ctx)

	return videoAdded, err
}

// VideoExists checks if a video with the given ID exists in the database.
func VideoExists(ctx context.Context, id string) bool {
	_, err := Client.Video.
		Query().
		Where(video.ID(id)).
		Only(ctx)

	return err == nil
}

// GetVideoByV gets a video by its ID.
func GetVideoByV(ctx context.Context, id string) (*ent.Video, error) {
	videoQueried, err := Client.Video.
		Query().
		Where(video.IDEQ(id)).
		Only(ctx)
	return videoQueried, err
}
