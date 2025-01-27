package models

import (
	"time"
)

// DeviceList represents a collection of devices associated with a user
type DeviceList struct {
	CurrentDeviceID uint     `json:"current_device_id"`
	Devices         []Device `json:"devices"`
}

// Device represents an individual device associated with a user
type Device struct {
	ID         uint      `json:"id"`
	LastActive time.Time `json:"last_active"`
	UserAgent  string    `json:"user_agent"`
	IPAddress  string    `json:"ip_address"`
}
