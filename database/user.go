package database

import (
	"context"

	"github.com/google/uuid"
	"github.com/hackinggate/alttube-go/ent"
	"github.com/hackinggate/alttube-go/ent/refreshtoken"
	"github.com/hackinggate/alttube-go/ent/user"
	"github.com/hackinggate/alttube-go/models"
)

// DeleteUserByID deletes a user by ID.
func DeleteUserByID(ctx context.Context, id uuid.UUID) error {
	err := Client.User.
		DeleteOneID(id).
		Exec(ctx)
	return err
}

// GetUserByID gets a user by ID.
func GetUserByID(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	userQueried, err := Client.User.
		Query().
		Where(user.IDEQ(id)).
		Only(ctx)
	return userQueried, err
}

// GetUserByRefreshToken gets a user by refresh token.
func GetUserByRefreshToken(ctx context.Context, token string) (*ent.User, error) {
	userQueried, err := Client.User.
		Query().
		Where(
			user.HasRefreshTokensWith(refreshtoken.Token(token)),
		).
		Only(ctx)
	return userQueried, err
}

// AddUser adds a new user to the database.
func AddUser(ctx context.Context, userToAdd models.SignupRequest) (*ent.User, error) {
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
func UpdateUserByID(ctx context.Context, id uuid.UUID, updateEmailRequest models.UpdateEmailRequest) error {
	_, err := Client.User.
		UpdateOneID(id).
		SetEmail(updateEmailRequest.Email).
		Save(ctx)
	return err
}
