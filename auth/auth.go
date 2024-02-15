package auth

import (
	"AltTube-Go/models"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

var jwtKey = []byte("my_secret_key")
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

var refreshTokens map[string]string // Maps refresh token to UUID

func init() {
	refreshTokens = make(map[string]string)
}

func AddRefreshToken(token, uuid string) {
	refreshTokens[token] = uuid
}

func ValidateRefreshToken(token string) (string, bool) {
	uuid, exists := refreshTokens[token]
	return uuid, exists
}

func RemoveRefreshToken(token string) {
	delete(refreshTokens, token)
}

// Generate Access Token
func GenerateAccessToken(uuid string) (string, error) {
	return generateToken(uuid, "access", 5*time.Minute) // Shorter expiration for access token
}

// Generate Refresh Token
func GenerateRefreshToken(uuid string) (string, error) {
	return generateToken(uuid, "refresh", 24*time.Hour) // Longer expiration for refresh token
}

// Helper function to generate tokens
func generateToken(uuid string, tokenType string, expiration time.Duration) (string, error) {
	expirationTime := time.Now().Add(expiration)
	claims := &models.Claims{
		UUID:      uuid,
		TokenType: tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
