package models

import (
	"time"
)

type DeviceListResponse struct {
	CurrentDeviceID uint             `json:"current_device_id"`
	Devices         []DeviceResponse `json:"devices"`
}

type DeviceResponse struct {
	ID         uint      `json:"id"`
	LastActive time.Time `json:"last_active"`
	UserAgent  string    `json:"user_agent"`
	IPAddress  string    `json:"ip_address"`
}
