package database

import (
	"AltTube-Go/models"
	"database/sql"
	"errors"

	"github.com/Masterminds/squirrel"
)

// AddLikeVideo creates a new like record in the database.
func AddLikeVideo(user *models.User, video *models.Video) error {
	query := dbBuilder.Insert("like_videos").
		Columns("user_id", "video_id").
		Values(user.ID, video.ID)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, execErr := dbInstance.Exec(sqlQuery, args...)
	return execErr
}

// ReadIsLikedVideo checks if a like record exists for a given user and video.
func ReadIsLikedVideo(user *models.User, video *models.Video) (bool, error) {
	query := dbBuilder.Select("1").
		From("like_videos").
		Where(squirrel.And{
			squirrel.Eq{"user_id": user.ID},
			squirrel.Eq{"video_id": video.ID},
		}).
		Limit(1)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return false, err
	}

	var exists int
	err = dbInstance.QueryRow(sqlQuery, args...).Scan(&exists)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

// RemoveLikeVideo deletes a like record for a given user and video.
func RemoveLikeVideo(user *models.User, video *models.Video) error {
	query := dbBuilder.Delete("like_videos").
		Where(squirrel.And{
			squirrel.Eq{"user_id": user.ID},
			squirrel.Eq{"video_id": video.ID},
		})

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, execErr := dbInstance.Exec(sqlQuery, args...)
	return execErr
}

// RemoveAllLikesByUserID deletes all like records for a given user.
func RemoveAllLikesByUserID(userID string) error {
	query := dbBuilder.Delete("like_videos").
		Where(squirrel.Eq{"user_id": userID})

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, execErr := dbInstance.Exec(sqlQuery, args...)
	return execErr
}

// GetAllLikesByUserID retrieves all like records for a given user.
func GetAllLikesByUserID(userID string) ([]models.LikeVideo, error) {
	query := dbBuilder.Select("id", "user_id", "video_id", "created_at", "updated_at").
		From("like_videos").
		Where(squirrel.Eq{"user_id": userID})

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := dbInstance.Query(sqlQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var likes []models.LikeVideo
	for rows.Next() {
		var like models.LikeVideo
		if err := rows.Scan(&like.ID, &like.UserID, &like.VideoID, &like.CreatedAt, &like.UpdatedAt); err != nil {
			return nil, err
		}
		likes = append(likes, like)
	}
	return likes, nil
}
