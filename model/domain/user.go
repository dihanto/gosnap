package domain

import "github.com/google/uuid"

type User struct {
	Id             uuid.UUID
	Username       string
	Name           string
	Email          string
	Password       string
	ProfilePicture string
	CreatedAt      int32
	UpdatedAt      int32
	DeletedAt      int32
}
