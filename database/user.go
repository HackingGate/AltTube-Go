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

// UpdateUserByID updates a user by ID.
func UpdateUserByID(ctx context.Context, id string, updateEmailRequest dto.UpdateEmailRequest) error {
	_, err := Client.User.
		UpdateOneID(id).
		SetEmail(updateEmailRequest.Email).
		Save(ctx)
	return err
}
