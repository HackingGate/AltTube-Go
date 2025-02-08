package auth

import (
	"github.com/google/uuid"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

type tokenClaims struct {
	UserID    uuid.UUID `json:"UserID"`
	TokenType string    `json:"token_type"`
	jwt.RegisteredClaims
}
