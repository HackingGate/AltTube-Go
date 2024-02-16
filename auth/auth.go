package auth

import (
	"AltTube-Go/models"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strings"
	"time"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))
var tokens []string

func AddToken(token string) {
	tokens = append(tokens, token)
}

func RemoveToken(tokenString string) {
	// Normalize tokenString by removing potential "Bearer " prefix
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	for i, t := range tokens {
		if t == tokenString {
			tokens = append(tokens[:i], tokens[i+1:]...)
			break
		}
	}
}

// GenerateAccessToken Generate access token with a short expiration
func GenerateAccessToken(uuid string) (string, error) {
	expiration := 5 * time.Minute // Short expiration
	token, _, err := generateToken(uuid, "access", expiration)
	return token, err
}

// GenerateRefreshToken Generate refresh token with a longer expiration
func GenerateRefreshToken(uuid string) (string, time.Time, error) {
	expiration := 24 * 30 * time.Hour // Longer expiration
	token, expiry, err := generateToken(uuid, "refresh", expiration)
	return token, expiry, err
}

// Unified token generation function
func generateToken(uuid string, tokenType string, expiration time.Duration) (string, time.Time, error) {
	expirationTime := time.Now().Add(expiration)
	claims := &models.Claims{
		UUID:      uuid,
		TokenType: tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	return tokenString, expirationTime, err
}
