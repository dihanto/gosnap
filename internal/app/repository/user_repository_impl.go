package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/dihanto/gosnap/internal/app/helper"
	"github.com/dihanto/gosnap/model/domain"
	"github.com/google/uuid"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

// UserRegister is a method to register a new user in the database.
func (repository *UserRepositoryImpl) UserRegister(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	user.CreatedAt = int32(time.Now().Unix())

	password, err := helper.HashPassword(user.Password)
	if err != nil {
		return domain.User{}, err
	}

	query := "INSERT INTO users (id, username, email, password, age, created_at) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err = tx.ExecContext(ctx, query, user.Id, user.Username, user.Email, password, user.Age, user.CreatedAt)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

// UserLogin is a method to authenticate a user during login.
func (repository *UserRepositoryImpl) UserLogin(ctx context.Context, tx *sql.Tx, username string, password string) (bool, uuid.UUID, error) {
	var pwd string
	var id uuid.UUID
	query := "SELECT password, id FROM users WHERE username = $1;"
	err := tx.QueryRowContext(ctx, query, username).Scan(&pwd, &id)
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
func (repository *UserRepositoryImpl) UserUpdate(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	user.UpdatedAt = int32(time.Now().Unix())
	query := "UPDATE users SET email=$1, username=$2, updated_at=$3 WHERE id=$4"
	_, err := tx.ExecContext(ctx, query, user.Email, user.Username, user.UpdatedAt, user.Id)
	if err != nil {
		return domain.User{}, err
	}

	queryAge := "SELECT age FROM users WHERE id=$1"
	rows, err := tx.QueryContext(ctx, queryAge, user.Id)
	if err != nil {
		return domain.User{}, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&user.Age)
		if err != nil {
			return domain.User{}, err
		}
	}

	return user, nil
}

// UserDelete is a method to "soft delete" a user in the database by setting the deleted_at field.
func (repository *UserRepositoryImpl) UserDelete(ctx context.Context, tx *sql.Tx, id uuid.UUID) error {
	deleteTime := int32(time.Now().Unix())

	query := "UPDATE users SET deleted_at=$1 WHERE id=$2"
	_, err := tx.ExecContext(ctx, query, deleteTime, id)
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}
