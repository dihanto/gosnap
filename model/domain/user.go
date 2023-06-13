package domain

import "github.com/google/uuid"

type User struct {
	Id        uuid.UUID
	Username  string
	Email     string
	Password  string
	Age       int
	CreatedAt int32
	UpdatedAt int32
	DeletedAt int32
}
