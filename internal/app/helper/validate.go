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

	query := "SELECT username FROM users"
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
	photoId := field.Param()
	var likeId int

	conn, _ := config.InitDatabaseConnection()
	defer conn.Close()

	queryLikes := "SELECT id FROM likes WHERE photo_id=$1"
	err := conn.QueryRowContext(context.Background(), queryLikes, photoId).Scan(&likeId)
	if err != nil {
		log.Fatal(err)
	}

	queryLikeDetail := "SELECT user_id FROM like_details WHERE like_id=$1"
	rows, err := conn.QueryContext(context.Background(), queryLikeDetail, likeId)
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

func ValidateUserNotFollowTwice(field validator.FieldLevel) bool {
	value := field.Field().Interface().(string)
	targetUsername := field.Param()
	var followId int

	conn, _ := config.InitDatabaseConnection()
	defer conn.Close()

	queryFollowers := "SELECT id FROM followers WHERE username=$1"
	err := conn.QueryRowContext(context.Background(), queryFollowers, targetUsername).Scan(&followId)
	if err != nil {
		log.Fatalln(err)
	}

	queryFollowerDetails := "SELECT follower_name FROM follower_details WHERE follow_id=$1"
	rows, err := conn.QueryContext(context.Background(), queryFollowerDetails, followId)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	var followerUsername string
	for rows.Next() {
		err = rows.Scan(&followerUsername)
		if err != nil {
			log.Fatalln(err)
		}
		if value == followerUsername {
			return false
		}
	}

	return true
}
