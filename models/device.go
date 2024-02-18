package models

import (
	"time"
)

type Device struct {
	ID         uint      `json:"id"`
	LastActive time.Time `json:"last_active"`
	UserAgent  string    `json:"user_agent"`
	IPAddress  string    `json:"ip_address"`
}
