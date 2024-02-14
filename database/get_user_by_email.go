package database

import "AltTube-Go/models"

func GetUserByEmail(email string) (*models.User, error) {
	// Query user by email
	result := models.User{}
	dbResult := dbInstance.Where("email = ?", email).First(&result)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return &result, nil
}
