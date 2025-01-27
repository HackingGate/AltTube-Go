package database

import (
	"AltTube-Go/models"
	"github.com/Masterminds/squirrel"
	"time"
)

var dbBuilder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

// AddAccessToken creates and stores a new access token in the database.
func AddAccessToken(at models.AccessToken) error {
	query := dbBuilder.Insert("access_tokens").
		Columns("token", "user_id", "expiry", "refresh_token_id", "created_at", "updated_at").
		Values(at.Token, at.UserID, at.Expiry, at.RefreshTokenID, at.CreatedAt, at.UpdatedAt)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, execErr := dbInstance.Exec(sqlQuery, args...)
	return execErr
}

// ValidateAccessToken checks if the given token exists and is not expired.
func ValidateAccessToken(token string) (string, error) {
	query := dbBuilder.Select("user_id").
		From("access_tokens").
		Where(squirrel.And{
			squirrel.Eq{"token": token},
			squirrel.Gt{"expiry": time.Now()},
		}).
		Limit(1)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return "", err
	}
	var userID string
	err = dbInstance.QueryRow(sqlQuery, args...).Scan(&userID)
	if err != nil {
		return "", err
	}
	return userID, nil
}

// RemoveAccessToken deletes an access token from the database.
func RemoveAccessToken(token string) error {
	query := dbBuilder.Delete("access_tokens").
		Where(squirrel.Eq{"token": token})

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, execErr := dbInstance.Exec(sqlQuery, args...)
	return execErr
}

// RemoveAllAccessTokensByRefreshTokenID deletes all access tokens associated with a specific refresh token.
func RemoveAllAccessTokensByRefreshTokenID(refreshTokenID uint) error {
	query := dbBuilder.Delete("access_tokens").
		Where(squirrel.Eq{"refresh_token_id": refreshTokenID})

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, execErr := dbInstance.Exec(sqlQuery, args...)
	return execErr
}

// AddRefreshToken creates and stores a new refresh token in the database.
func AddRefreshToken(rt models.RefreshToken) error {
	query := dbBuilder.Insert("refresh_tokens").
		Columns("token", "user_id", "expiry", "user_agent", "ip_address", "created_at", "updated_at").
		Values(rt.Token, rt.UserID, rt.Expiry, rt.UserAgent, rt.IPAddress, rt.CreatedAt, rt.UpdatedAt)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, execErr := dbInstance.Exec(sqlQuery, args...)
	return execErr
}

// ValidateRefreshToken checks if the given refresh token exists and is not expired.
func ValidateRefreshToken(token string) (string, error) {
	query := dbBuilder.Select("user_id").
		From("refresh_tokens").
		Where(squirrel.And{
			squirrel.Eq{"token": token},
			squirrel.Gt{"expiry": time.Now()},
		}).
		Limit(1)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return "", err
	}
	var userID string
	err = dbInstance.QueryRow(sqlQuery, args...).Scan(&userID)
	if err != nil {
		return "", err
	}
	return userID, nil
}

// RemoveRefreshTokenByToken deletes a refresh token from the database.
func RemoveRefreshTokenByToken(token string) error {
	query := dbBuilder.Delete("refresh_tokens").
		Where(squirrel.Eq{"token": token})

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, execErr := dbInstance.Exec(sqlQuery, args...)
	return execErr
}

// RemoveRefreshTokensByID deletes multiple refresh tokens by their IDs.
func RemoveRefreshTokensByID(ids []uint) error {
	query := dbBuilder.Delete("refresh_tokens").
		Where(squirrel.Eq{"id": ids})

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, execErr := dbInstance.Exec(sqlQuery, args...)
	return execErr
}

// GetAllRefreshTokensByUserID retrieves all refresh tokens for a specific user ID.
func GetAllRefreshTokensByUserID(userID string) ([]models.RefreshToken, error) {
	query := dbBuilder.Select("id", "token", "user_id", "expiry", "user_agent", "ip_address", "created_at", "updated_at").
		From("refresh_tokens").
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

	var refreshTokens []models.RefreshToken
	for rows.Next() {
		var rt models.RefreshToken
		if err := rows.Scan(&rt.ID, &rt.Token, &rt.UserID, &rt.Expiry, &rt.UserAgent, &rt.IPAddress, &rt.CreatedAt, &rt.UpdatedAt); err != nil {
			return nil, err
		}
		refreshTokens = append(refreshTokens, rt)
	}
	return refreshTokens, nil
}

// GetRefreshTokenByToken retrieves a refresh token by its token value.
func GetRefreshTokenByToken(token string) (models.RefreshToken, error) {
	query := dbBuilder.Select("id", "token", "user_id", "expiry", "user_agent", "ip_address", "created_at", "updated_at").
		From("refresh_tokens").
		Where(squirrel.Eq{"token": token}).
		Limit(1)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return models.RefreshToken{}, err
	}

	var rt models.RefreshToken
	err = dbInstance.QueryRow(sqlQuery, args...).Scan(&rt.ID, &rt.Token, &rt.UserID, &rt.Expiry, &rt.UserAgent, &rt.IPAddress, &rt.CreatedAt, &rt.UpdatedAt)
	if err != nil {
		return models.RefreshToken{}, err
	}
	return rt, nil
}

// GetRefreshTokenByAccessToken retrieves a refresh token associated with a specific access token.
func GetRefreshTokenByAccessToken(accessToken string) (models.RefreshToken, error) {
	query := dbBuilder.Select("refresh_tokens.id", "refresh_tokens.token", "refresh_tokens.user_id", "refresh_tokens.expiry", "refresh_tokens.user_agent", "refresh_tokens.ip_address", "refresh_tokens.created_at", "refresh_tokens.updated_at").
		From("refresh_tokens").
		Join("access_tokens ON access_tokens.refresh_token_id = refresh_tokens.id").
		Where(squirrel.Eq{"access_tokens.token": accessToken}).
		Limit(1)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return models.RefreshToken{}, err
	}

	var rt models.RefreshToken
	err = dbInstance.QueryRow(sqlQuery, args...).Scan(&rt.ID, &rt.Token, &rt.UserID, &rt.Expiry, &rt.UserAgent, &rt.IPAddress, &rt.CreatedAt, &rt.UpdatedAt)
	if err != nil {
		return models.RefreshToken{}, err
	}
	return rt, nil
}
