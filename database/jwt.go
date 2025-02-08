package database

import (
	"AltTube-Go/ent"
	"AltTube-Go/ent/accesstoken"
	"AltTube-Go/ent/refreshtoken"
	"AltTube-Go/ent/user"
	"context"
	"time"
)

// AddAccessToken creates and stores a new access token in the database.
func AddAccessToken(ctx context.Context, token string, user *ent.User, expiry time.Time, refreshToken *ent.RefreshToken) error {
	// Create a new access token entity.
	_, err := Client.AccessToken.
		Create().
		SetToken(token).
		SetExpiry(expiry).
		SetUser(user).
		SetRefreshToken(refreshToken).
		Save(ctx)
	return err
}

// AddRefreshToken creates and stores a new refresh token in the database.
func AddRefreshToken(ctx context.Context, token string, user *ent.User, expiry time.Time, userAgent string, ipAddress string) error {
	// Create a new refresh token entity.
	_, err := Client.RefreshToken.
		Create().
		SetToken(token).
		SetExpiry(expiry).
		SetUser(user).
		SetUserAgent(userAgent).
		SetIPAddress(ipAddress).
		Save(ctx)
	return err
}

// RemoveAccessTokenByRefreshToken deletes all access tokens associated with the specified refresh token.
func RemoveAccessTokenByRefreshToken(ctx context.Context, token string) error {
	// Query the database for the refresh token with the given token.
	refreshToken, err := Client.RefreshToken.
		Query().
		Where(
			refreshtoken.Token(token),
		).
		Only(ctx)

	// If an error occurs or no result is found, return the error.
	if err != nil {
		return err
	}

	// Delete all access tokens associated with the refresh token.
	_, err = Client.AccessToken.
		Delete().
		Where(
			accesstoken.HasRefreshTokenWith(refreshtoken.IDEQ(refreshToken.ID)),
		).
		Exec(ctx)
	return err
}

// RemoveRefreshTokenByToken deletes the refresh token with the given token.
func RemoveRefreshTokenByToken(ctx context.Context, token string) error {
	// Delete the refresh token with the given token.
	_, err := Client.RefreshToken.
		Delete().
		Where(
			refreshtoken.Token(token),
		).
		Exec(ctx)
	return err
}

// ValidateRefreshToken checks if the given token exists and is not expired.
func ValidateRefreshToken(ctx context.Context, tokenString string) (string, bool) {
	// Query database for the token
	refreshToken, err := Client.RefreshToken.
		Query().
		Where(
			refreshtoken.Token(tokenString),
			refreshtoken.ExpiryGT(time.Now()), // Ensure the token is still valid
		).
		Only(ctx)

	// If an error occurs or no result is found, return false
	if err != nil {
		return "", false
	}

	// Return token and success status
	return refreshToken.Token, true
}

// ValidateAccessToken checks if the given token exists and is not expired.
func ValidateAccessToken(ctx context.Context, tokenString string) (string, bool) {
	// Query database for the token
	accessToken, err := Client.AccessToken.
		Query().
		Where(
			accesstoken.Token(tokenString),
			accesstoken.ExpiryGT(time.Now()), // Ensure the token is still valid
		).
		Only(ctx)

	// If an error occurs or no result is found, return false
	if err != nil {
		return "", false
	}

	// Return user ID and success status
	return accessToken.UserID, true
}

// RemoveAccessToken deletes the access token with the given token.
func RemoveAccessToken(ctx context.Context, token string) error {
	// Delete the access token with the given token.
	_, err := Client.AccessToken.
		Delete().
		Where(
			accesstoken.Token(token),
		).
		Exec(ctx)
	return err
}

// GetAllRefreshTokensByUserID returns all refresh tokens associated with the given user.
func GetAllRefreshTokensByUserID(ctx context.Context, userID string) ([]*ent.RefreshToken, error) {
	// Query the database for refresh tokens belonging to the user.
	refreshTokens, err := Client.RefreshToken.
		Query().
		Where(
			refreshtoken.HasUserWith(user.IDEQ(userID)),
		).
		All(ctx)
	return refreshTokens, err
}

// RemoveAllAccessTokensByRefreshTokenID deletes all access tokens associated with the specified refresh token.
func RemoveAllAccessTokensByRefreshTokenID(ctx context.Context, refreshTokenID uint) error {
	// Delete access tokens where the refresh token ID matches.
	_, err := Client.AccessToken.
		Delete().
		Where(
			accesstoken.HasRefreshTokenWith(refreshtoken.IDEQ(refreshTokenID)),
		).
		Exec(ctx)
	return err
}

// RemoveRefreshTokensByID deletes all refresh tokens with IDs contained in the provided slice.
func RemoveRefreshTokensByID(ctx context.Context, ids []uint) error {
	// Delete refresh tokens in batch using their IDs.
	_, err := Client.RefreshToken.
		Delete().
		Where(
			refreshtoken.IDIn(ids...),
		).
		Exec(ctx)
	return err
}

// GetRefreshTokenByAccessToken retrieves the refresh token associated with the given access token.
func GetRefreshTokenByAccessToken(ctx context.Context, accessToken string) (*ent.RefreshToken, error) {
	// Query the database for the refresh token associated with the access token.
	refreshToken, err := Client.RefreshToken.
		Query().
		Where(
			refreshtoken.HasAccessTokensWith(accesstoken.Token(accessToken)),
		).
		Only(ctx)
	return refreshToken, err
}

// GetRefreshTokenByToken retrieves the refresh token with the given token.
func GetRefreshTokenByToken(ctx context.Context, token string) (*ent.RefreshToken, error) {
	// Query the database for the refresh token with the given token.
	refreshToken, err := Client.RefreshToken.
		Query().
		Where(
			refreshtoken.Token(token),
		).
		Only(ctx)
	return refreshToken, err
}
