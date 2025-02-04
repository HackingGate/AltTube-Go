package database

import (
	"AltTube-Go/dto"
	"AltTube-Go/ent"
	"AltTube-Go/ent/user"
	"context"
)

// DeleteUserByID deletes a user by ID.
func DeleteUserByID(ctx context.Context, id string) error {
	err := Client.User.
		DeleteOneID(id).
		Exec(ctx)
	return err
}

// GetUserByID gets a user by ID.
func GetUserByID(ctx context.Context, id string) (*ent.User, error) {
	userQueried, err := Client.User.
		Query().
		Where(user.IDEQ(id)).
		Only(ctx)
	return userQueried, err
}

// AddUser adds a new user to the database.
func AddUser(ctx context.Context, userToAdd dto.SignupRequest) (*ent.User, error) {
	userAdded, err := Client.User.
		Create().
		SetEmail(userToAdd.Email).
		SetPassword(userToAdd.Password).
		Save(ctx)
	return userAdded, err
}

// GetUserByEmail gets a user by email.
func GetUserByEmail(ctx context.Context, email string) (*ent.User, error) {
	userQueried, err := Client.User.
		Query().
		Where(user.EmailEQ(email)).
		Only(ctx)
	return userQueried, err
}

//func GetUserByEmail(email string) (*models.User, error) {
//	// Query user by email
//	result := models.User{}
//	dbResult := dbInstance.Where("email = ?", email).First(&result)
//	if dbResult.Error != nil {
//		return nil, dbResult.Error
//	}
//	return &result, nil
//}

// UpdateUserByID updates a user by ID.
func UpdateUserByID(ctx context.Context, id string, updateEmailRequest dto.UpdateEmailRequest) error {
	_, err := Client.User.
		UpdateOneID(id).
		SetEmail(updateEmailRequest.Email).
		Save(ctx)
	return err
}

//func UpdateUserByID(id string, editEmail models.EditEmail) error {
//	// Check if the user exists
//	var existingUser models.User
//	dbResult := dbInstance.Where("id = ?", id).First(&existingUser)
//	if dbResult.Error != nil {
//		if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
//			// User does not exist
//			return errors.New("user not found")
//		}
//		// Other error
//		return dbResult.Error
//	}
//
//	// Perform the update
//	dbResult = dbInstance.Model(&existingUser).Updates(editEmail)
//	if dbResult.Error != nil {
//		return dbResult.Error
//	}
//	return nil
//}
//
//func DeleteUserByID(id string) error {
//	// Perform a hard delete (completely remove) the user by id
//	dbResult := dbInstance.Unscoped().Where("id = ?", id).Delete(&models.User{})
//	if dbResult.Error != nil {
//		return dbResult.Error
//	}
//	return nil
//}
