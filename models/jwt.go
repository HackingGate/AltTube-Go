package models

import (
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID    string `json:"UserID"`
	TokenType string `json:"token_type"`
	jwt.RegisteredClaims
}
