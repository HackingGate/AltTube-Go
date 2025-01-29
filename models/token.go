package models

import (
	"time"
)

// AccessToken represents an access token in the database
type AccessToken struct {
	Entity
	Token          string    `json:"token"`
	UserID         string    `json:"user_id"`
	Expiry         time.Time `json:"expiry"`
	RefreshTokenID uint      `json:"refresh_token_id"`
}

// RefreshToken represents a refresh token in the database
type RefreshToken struct {
	Entity
	Token     string    `json:"token"`
	Expiry    time.Time `json:"expiry"`
	UserAgent string    `json:"user_agent"`
	IPAddress string    `json:"ip_address"`
	UserID    string    `json:"user_id"`
}
