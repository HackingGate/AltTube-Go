package models

import (
	"github.com/golang-jwt/jwt/v5"
)

// Claims represents the JWT claims for a token
type Claims struct {
	UserID    string `json:"UserID"`
	TokenType string `json:"token_type"`
	jwt.RegisteredClaims
}
