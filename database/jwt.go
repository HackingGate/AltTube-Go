package database

import (
	"AltTube-Go/models"
	"time"
)

// AddAccessToken creates and stores a new access token in the database.
func AddAccessToken(token, userID string, expiry time.Time, refreshToken models.RefreshToken) error {
	accessToken := models.AccessToken{
		Token:          token,
		UserID:         userID,
		Expiry:         expiry,
		RefreshTokenID: refreshToken.ID,
		RefreshToken:   refreshToken,
	}
	return dbInstance.Create(&accessToken).Error
}

// ValidateAccessToken checks if the given token exists and is not expired.
func ValidateAccessToken(token string) (string, bool) {
	var accessToken models.AccessToken
	result := dbInstance.Where("token = ? AND expiry > ?", token, time.Now()).First(&accessToken)
	if result.Error != nil || result.RowsAffected == 0 {
		return "", false // Token not found or expired
	}
	return accessToken.UserID, true
}

// RemoveAccessToken deletes an access token from the database.
func RemoveAccessToken(token string) error {
	result := dbInstance.Unscoped().Where("token = ?", token).Delete(&models.AccessToken{})
	return result.Error
}

// GetAllAccessTokensByUserID returns all access tokens for a given user.
func GetAllAccessTokensByUserID(userID string) ([]string, error) {
	var accessTokens []models.AccessToken // Use a slice to hold multiple tokens
	result := dbInstance.Where("user_id = ? AND expiry > ?", userID, time.Now()).Find(&accessTokens)

	if result.Error != nil {
		return nil, result.Error
	}

	// Extract the token strings from the accessTokens slice
	var tokens []string
	for _, accessToken := range accessTokens {
		tokens = append(tokens, accessToken.Token)
	}

	return tokens, nil
}

func RemoveAllAccessTokensByRefreshTokenID(refreshTokenID uint) error {
	result := dbInstance.Unscoped().Where("refresh_token_id = ?", refreshTokenID).Delete(&models.AccessToken{})
	return result.Error
}

// AddRefreshToken creates and stores a new refresh token in the database.
func AddRefreshToken(token, userID string, expiry time.Time, userAgent string, ipAddress string) error {
	refreshToken := models.RefreshToken{
		Token:     token,
		UserID:    userID,
		Expiry:    expiry,
		UserAgent: userAgent,
		IPAddress: ipAddress,
	}
	return dbInstance.Create(&refreshToken).Error
}

// ValidateRefreshToken checks if the given token exists and is not expired.
func ValidateRefreshToken(token string) (string, bool) {
	var refreshToken models.RefreshToken
	result := dbInstance.Where("token = ? AND expiry > ?", token, time.Now()).First(&refreshToken)
	if result.Error != nil || result.RowsAffected == 0 {
		return "", false // Token not found or expired
	}
	return refreshToken.UserID, true
}

// RemoveRefreshToken deletes a refresh token from the database.
func RemoveRefreshToken(token string) error {
	// First, get the refresh token to get its ID
	refreshToken, err := GetRefreshTokenByToken(token)
	if err != nil {
		return err
	}

	// Delete all access tokens associated with the refresh token
	err = RemoveAllAccessTokensByRefreshTokenID(refreshToken.ID)
	if err != nil {
		return err
	}

	result := dbInstance.Unscoped().Where("token = ?", token).Delete(&models.RefreshToken{})
	return result.Error
}

func GetAllRefreshTokensByUserID(userID string) ([]string, error) {
	var refreshTokens []models.RefreshToken // Use a slice to hold multiple tokens
	result := dbInstance.Where("user_id = ? AND expiry > ?", userID, time.Now()).Find(&refreshTokens)

	if result.Error != nil {
		return nil, result.Error
	}

	// Extract the token strings from the refreshTokens slice
	var tokens []string
	for _, refreshToken := range refreshTokens {
		tokens = append(tokens, refreshToken.Token)
	}

	return tokens, nil
}

func GetRefreshTokenByToken(token string) (models.RefreshToken, error) {
	var refreshToken models.RefreshToken
	result := dbInstance.Where("token = ?", token).First(&refreshToken)
	return refreshToken, result.Error
}
