package model

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	v string `gorm:"primary unique"`
}
