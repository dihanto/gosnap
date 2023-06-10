package helper

import (
	"context"

	"github.com/dihanto/gosnap/internal/app/config"
	"github.com/go-playground/validator/v10"
)

func ValidateEmailUniq(field validator.FieldLevel) bool {
	value := field.Field().Interface().(string)

	conn := config.NewDb()
	defer conn.Close()

	ctx := context.Background()
	query := "select email from users"
	rows, err := conn.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var email string
		err = rows.Scan(&email)
		if err != nil {
			panic(err)
		}
		if value == email {
			return false
		}
	}

	return true

}

func ValidateUsernameUniq(field validator.FieldLevel) bool {
	value := field.Field().Interface().(string)

	conn := config.NewDb()
	defer conn.Close()

	query := "select username from users"
	ctx := context.Background()
	rows, err := conn.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var username string
		err = rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		if value == username {
			return false
		}
	}
	return true
}
