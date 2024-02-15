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

func GenerateJWT(uuid string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &models.Claims{
		UUID: uuid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)
}

// contains checks if the tokens slice contains a specific token.
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
