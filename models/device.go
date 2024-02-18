package models

import (
	"time"
)

type Device struct {
	ID         uint
	LastActive time.Time
	UserAgent  string
	IPAddress  string
}
