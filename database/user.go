package database

import (
	"AltTube-Go/models"
	"errors"
	"gorm.io/gorm"
)

func AddUser(user models.User) error {
	// Check if the user already exists
	existingUser, err := GetUserByEmail(user.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		// An error occurred that isn't the record not found error
		return err
	}
	if existingUser != nil {
		// User already exists
		return errors.New("user already exists")
	}

	// If user does not exist, create the user
	dbResult := dbInstance.Create(&user)
	return dbResult.Error
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

func UpdateUserByUUID(uuid string, editEmail models.EditEmail) error {
	// Check if the user exists
	var existingUser models.User
	dbResult := dbInstance.Where("uuid = ?", uuid).First(&existingUser)
	if dbResult.Error != nil {
		if dbResult.Error == gorm.ErrRecordNotFound {
			// User does not exist
			return errors.New("user not found")
		}
		// Other error
		return dbResult.Error
	}

	// Perform the update
	dbResult = dbInstance.Model(&existingUser).Updates(editEmail)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}

func DeleteUserByUUID(uuid string) error {
	// Perform a hard delete (completely remove) the user by uuid
	dbResult := dbInstance.Unscoped().Where("uuid = ?", uuid).Delete(&models.User{})
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}
