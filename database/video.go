package database

import (
	"AltTube-Go/models"
	"database/sql"

	"github.com/Masterminds/squirrel"
)

// AddVideo inserts a new video into the database.
func AddVideo(video models.Video) error {
	query := dbBuilder.Insert("videos").
		Columns("id", "title", "description", "created_at", "updated_at").
		Values(video.ID, video.Title, video.Description, video.CreatedAt, video.UpdatedAt)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, execErr := dbInstance.Exec(sqlQuery, args...)
	return execErr
}

// GetVideoByV retrieves a video by its ID.
func GetVideoByV(id string) (*models.Video, error) {
	query := dbBuilder.Select("id", "title", "description", "created_at", "updated_at").
		From("videos").
		Where(squirrel.Eq{"id": id}).
		Limit(1)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	var video models.Video
	err = dbInstance.QueryRow(sqlQuery, args...).Scan(&video.ID, &video.Title, &video.Description, &video.CreatedAt, &video.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &video, nil
}

// GetUserByID retrieves a user by their ID.
func GetUserByID(id string) (*models.User, error) {
	query := dbBuilder.Select("id", "email", "created_at", "updated_at").
		From("users").
		Where(squirrel.Eq{"id": id}).
		Limit(1)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	var user models.User
	err = dbInstance.QueryRow(sqlQuery, args...).Scan(&user.ID, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// VideoExists checks if a video exists in the database by its ID.
func VideoExists(id string) bool {
	query := dbBuilder.Select("1").
		From("videos").
		Where(squirrel.Eq{"id": id}).
		Limit(1)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return false
	}

	var exists int
	err = dbInstance.QueryRow(sqlQuery, args...).Scan(&exists)
	return err == nil
}
