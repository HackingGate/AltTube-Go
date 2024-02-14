package database

import "AltTube-Go/models"

func AddUser(user models.User) error {
	dbInstance.Create(&user)
	return nil
}
