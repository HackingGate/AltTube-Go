package auth

import (
	"github.com/google/uuid"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateAccessToken Generate access token with a short expiration
func GenerateAccessToken(userID uuid.UUID) (string, time.Time, error) {
	expiration := 5 * time.Minute // Short expiration
	token, expiry, err := generateToken(userID, "access", expiration)
	return token, expiry, err
}

// GenerateRefreshToken Generate refresh token with a longer expiration
func GenerateRefreshToken(userID uuid.UUID) (string, time.Time, error) {
	expiration := 24 * 30 * time.Hour // Longer expiration
	token, expiry, err := generateToken(userID, "refresh", expiration)
	return token, expiry, err
}

// Unified token generation function
func generateToken(userID uuid.UUID, tokenType string, expiration time.Duration) (string, time.Time, error) {
	expirationTime := time.Now().Add(expiration)
	claims := &tokenClaims{
		UserID:    userID,
		TokenType: tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	return tokenString, expirationTime, err
}
