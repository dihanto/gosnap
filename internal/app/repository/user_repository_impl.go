package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/dihanto/gosnap/internal/app/helper"
	"github.com/dihanto/gosnap/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) UserRegister(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {

	t := time.Now()
	user.CreatedAt = int32(t.Unix())

	password, err := helper.HashPassword(user.Password)
	if err != nil {
		panic(err)
	}

	query := "INSERT INTO users (username, email, password, age, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	row := tx.QueryRowContext(ctx, query, user.Username, user.Email, password, user.Age, user.CreatedAt)

	var id int
	err = row.Scan(&id)
	if err != nil {
		return domain.User{}, err
	}

	user.Id = id

	return user, nil
}

func (repository *UserRepositoryImpl) UserLogin(ctx context.Context, tx *sql.Tx, username string, password string) (bool, int, error) {
	var pwd string
	var id int
	query := "SELECT password, id FROM users WHERE username = $1;"
	err := tx.QueryRowContext(ctx, query, username).Scan(&pwd, &id)
	if err == sql.ErrNoRows {
		fmt.Println("Usename not found")
		return false, 0, err
	}
	if err != nil {
		fmt.Println("query error")
		return false, 0, err
	}

	match, err := helper.CheckPasswordHash(password, pwd)
	if !match {
		fmt.Println("Password do not match")
		return false, 0, err
	}

	return true, id, nil

}

func (repository *UserRepositoryImpl) UserUpdate(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {

	t := time.Now()
	user.UpdatedAt = int32(t.Unix())
	query := "UPDATE users SET email=$1, username=$2, updated_at=$3 WHERE id=$4"
	rows, err := tx.QueryContext(ctx, query, user.Email, user.Username, user.UpdatedAt, user.Id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&user.Id, &user.Email, &user.Username, &user.Age, &user.UpdatedAt)
		if err != nil {
			panic(err)
		}
	}

	return user, nil
}

func (repository *UserRepositoryImpl) UserDelete(ctx context.Context, tx *sql.Tx, id int) error {

	t := time.Now()
	timestamp := int32(t.Unix())

	query := "UPDATE users SET deleted_at=$1 WHERE id=$2"
	_, err := tx.ExecContext(ctx, query, timestamp, id)
	if err != nil {
		panic(err)
	}
	return nil
}
