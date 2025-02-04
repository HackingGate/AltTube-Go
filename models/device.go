package models

import (
	"time"
)

type DeviceList struct {
	CurrentDeviceID uint     `json:"current_device_id"`
	Devices         []Device `json:"devices"`
}

type Device struct {
	ID         uint      `json:"id"`
	LastActive time.Time `json:"last_active"`
	UserAgent  string    `json:"user_agent"`
	IPAddress  string    `json:"ip_address"`
}
