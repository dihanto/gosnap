package domain

import "github.com/google/uuid"

type Comment struct {
	Id        int
	UserId    uuid.UUID
	PhotoId   int
	Message   string
	CreatedAt int32
	UpdatedAt int32
	DeletedAt int32
}
