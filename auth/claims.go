package auth

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

type tokenClaims struct {
	UserID    string `json:"UserID"`
	TokenType string `json:"token_type"`
	jwt.RegisteredClaims
}
