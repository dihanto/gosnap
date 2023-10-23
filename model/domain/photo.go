package domain

import "github.com/google/uuid"

type Photo struct {
	Id          int
	Title       string
	Caption     string
	PhotoBase64 string
	UserId      uuid.UUID
	CreatedAt   int32
	UpdatedAt   int32
	DeletedAt   int32
}
