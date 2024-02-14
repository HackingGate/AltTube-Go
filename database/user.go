package database

import "AltTube-Go/models"

func AddUser(user models.User) error {
	dbInstance.Create(&user)
	return nil
}

func GetUserByEmail(email string) (*models.User, error) {
	// Query user by email
	result := models.User{}
	dbResult := dbInstance.Where("email = ?", email).First(&result)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return &result, nil
}

func UpdateUserByEmail(email string, user models.User) error {
	// Update user by email
	dbResult := dbInstance.Model(&models.User{}).Where("email = ?", email).Updates(user)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}

func DeleteUserByEmail(email string) error {
	// Perform a hard delete (completely remove) the user by email
	dbResult := dbInstance.Unscoped().Where("email = ?", email).Delete(&models.User{})
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}
