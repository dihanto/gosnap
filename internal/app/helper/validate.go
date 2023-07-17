package helper

import (
	"context"
	"log"

	"github.com/dihanto/gosnap/internal/app/config"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func ValidateEmailUniq(field validator.FieldLevel) bool {
	value := field.Field().Interface().(string)

	conn, _ := config.InitDatabaseConnection()
	defer conn.Close()

	ctx := context.Background()
	query := "SELECT email FROM users"
	rows, err := conn.QueryContext(ctx, query)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var email string
		err = rows.Scan(&email)
		if err != nil {
			log.Println(err)
		}
		if value == email {
			return false
		}
	}

	return true

}

func ValidateUsernameUniq(field validator.FieldLevel) bool {
	value := field.Field().Interface().(string)

	conn, _ := config.InitDatabaseConnection()
	defer conn.Close()

	query := "SELECT username FFROM users"
	ctx := context.Background()
	rows, err := conn.QueryContext(ctx, query)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var username string
		err = rows.Scan(&username)
		if err != nil {
			log.Println(err)
		}
		if value == username {
			return false
		}
	}
	return true
}

func ValidateOneUserOneLike(field validator.FieldLevel) bool {
	value := field.Field().Interface().(uuid.UUID)
	id := field.Param()

	conn, _ := config.InitDatabaseConnection()
	defer conn.Close()

	query := "SELECT user_id FROM like_details WHERE photo_id=$1"
	rows, err := conn.QueryContext(context.Background(), query, id)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	var userId uuid.UUID
	for rows.Next() {
		err = rows.Scan(&userId)
		if err != nil {
			log.Fatalln(err)
		}
		if value == userId {
			return false
		}
	}

	return true
}
