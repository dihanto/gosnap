package repository

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
	"time"

	"github.com/dihanto/gosnap/internal/app/helper"
	"github.com/dihanto/gosnap/model/domain"
	"github.com/google/uuid"
)

type UserRepositoryImpl struct {
	Database *sql.DB
}

func NewUserRepository(database *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		Database: database,
	}
}

// UserRegister is a method to register a new user in the database.
func (repository *UserRepositoryImpl) UserRegister(ctx context.Context, user domain.User) (domain.User, error) {
	tx, err := repository.Database.Begin()
	if err != nil {
		return domain.User{}, err
	}
	defer helper.CommitOrRollback(tx, &err)

	user.CreatedAt = int32(time.Now().Unix())

	password, err := helper.HashPassword(user.Password)
	if err != nil {
		return domain.User{}, err
	}

	query := "INSERT INTO users (id, username, name, email, password, created_at) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err = tx.ExecContext(ctx, query, user.Id, user.Username, user.Name, user.Email, password, user.CreatedAt)
	if err != nil {
		return domain.User{}, err
	}

	queryFollow := "INSERT INTO followers (user_id, username) VALUES ($1, $2)"
	_, err = tx.ExecContext(ctx, queryFollow, user.Id, user.Username)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

// UserLogin is a method to authenticate a user during login.
func (repository *UserRepositoryImpl) UserLogin(ctx context.Context, username string, password string) (bool, uuid.UUID, error) {
	tx, err := repository.Database.Begin()
	if err != nil {
		return false, uuid.Nil, err
	}
	defer helper.CommitOrRollback(tx, &err)

	var pwd string
	var id uuid.UUID
	query := "SELECT password, id FROM users WHERE username = $1;"
	err = tx.QueryRowContext(ctx, query, username).Scan(&pwd, &id)
	if err == sql.ErrNoRows {
		return false, uuid.Nil, err
	}
	if err != nil {
		return false, uuid.Nil, err
	}

	match, err := helper.CheckPasswordHash(password, pwd)
	if !match {
		return false, uuid.Nil, err
	}

	return true, id, nil
}

// UserUpdate is a method to update user information in the database.
func (repository *UserRepositoryImpl) UserUpdate(ctx context.Context, user domain.User) (domain.User, error) {
	tx, err := repository.Database.Begin()
	if err != nil {
		return domain.User{}, nil
	}
	defer helper.CommitOrRollback(tx, &err)

	user.UpdatedAt = int32(time.Now().Unix())
	query := "UPDATE users SET"
	params := []interface{}{}
	paramCount := 1

	if user.Email != "" {
		query += " email = $" + strconv.Itoa(paramCount)
		params = append(params, user.Email)
		paramCount++
	}

	if user.Username != "" {
		if paramCount > 1 {
			query += ","
		}
		query += " username = $" + strconv.Itoa(paramCount)
		params = append(params, user.Username)
		paramCount++
	}

	if user.ProfilePicture != "" {
		if paramCount > 1 {
			query += ","
		}
		query += " profile_picture_base64 = $" + strconv.Itoa(paramCount)
		params = append(params, user.ProfilePicture)
		paramCount++
	}

	query += ", updated_at = $" + strconv.Itoa(paramCount)
	params = append(params, user.UpdatedAt)
	paramCount++

	query += " WHERE id = $" + strconv.Itoa(paramCount)
	params = append(params, user.Id)

	_, err = tx.ExecContext(ctx, query, params...)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

// UserDelete is a method to "soft delete" a user in the database by setting the deleted_at field.
func (repository *UserRepositoryImpl) UserDelete(ctx context.Context, id uuid.UUID) error {
	tx, err := repository.Database.Begin()
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx, &err)

	deleteTime := int32(time.Now().Unix())

	query := "UPDATE users SET deleted_at=$1 WHERE id=$2"
	_, err = tx.ExecContext(ctx, query, deleteTime, id)
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func (repository *UserRepositoryImpl) FindUser(ctx context.Context, id uuid.UUID) (user domain.User, err error) {
	tx, err := repository.Database.Begin()
	if err != nil {
		return
	}
	defer helper.CommitOrRollback(tx, &err)

	query := "SELECT username, name, profile_picture_base64 FROM users WHERE id=$1"
	err = tx.QueryRowContext(ctx, query, id).Scan(&user.Username, &user.Name, &user.ProfilePicture)
	if err != nil {
		return
	}

	return user, nil

}

func (repository *UserRepositoryImpl) FindAllUser(ctx context.Context) (users []domain.User, err error) {
	tx, err := repository.Database.Begin()
	if err != nil {
		return
	}
	defer helper.CommitOrRollback(tx, &err)

	query := "SELECT username, profile_picture_base64 FROM users"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user domain.User
		err = rows.Scan(&user.Username, &user.ProfilePicture)
		if err != nil {
			return
		}

		users = append(users, user)
	}

	return
}
