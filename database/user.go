package database

import (
	"AltTube-Go/models"
	"database/sql"
	"errors"

	"github.com/Masterminds/squirrel"
)

// AddUser inserts a new user into the database.
func AddUser(user models.User) error {
	// Check if the user already exists
	existingUser, err := GetUserByEmail(user.Email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}
	if existingUser != nil {
		return errors.New("user already exists")
	}

	// Insert the new user
	query := dbBuilder.Insert("users").
		Columns("id", "email", "created_at", "updated_at").
		Values(user.ID, user.Email, user.CreatedAt, user.UpdatedAt)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, execErr := dbInstance.Exec(sqlQuery, args...)
	return execErr
}

// GetUserByEmail retrieves a user by their email.
func GetUserByEmail(email string) (*models.User, error) {
	query := dbBuilder.Select("id", "email", "created_at", "updated_at").
		From("users").
		Where(squirrel.Eq{"email": email}).
		Limit(1)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	var user models.User
	err = dbInstance.QueryRow(sqlQuery, args...).Scan(&user.ID, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUserByID updates a user by their ID.
func UpdateUserByID(id string, editEmail models.EditEmail) error {
	// Check if the user exists
	queryCheck := dbBuilder.Select("id").
		From("users").
		Where(squirrel.Eq{"id": id}).
		Limit(1)

	sqlQueryCheck, argsCheck, err := queryCheck.ToSql()
	if err != nil {
		return err
	}

	var exists int
	err = dbInstance.QueryRow(sqlQueryCheck, argsCheck...).Scan(&exists)
	if errors.Is(err, sql.ErrNoRows) {
		return errors.New("user not found")
	}
	if err != nil {
		return err
	}

	// Update the user
	queryUpdate := dbBuilder.Update("users").
		Set("email", editEmail.Email).
		Where(squirrel.Eq{"id": id})

	sqlQueryUpdate, argsUpdate, err := queryUpdate.ToSql()
	if err != nil {
		return err
	}
	_, execErr := dbInstance.Exec(sqlQueryUpdate, argsUpdate...)
	return execErr
}

// DeleteUserByID deletes a user by their ID.
func DeleteUserByID(id string) error {
	query := dbBuilder.Delete("users").
		Where(squirrel.Eq{"id": id})

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, execErr := dbInstance.Exec(sqlQuery, args...)
	return execErr
}
