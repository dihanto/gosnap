package domain

import "github.com/google/uuid"

type Comment struct {
	Id        uuid.UUID
	UserId    uuid.UUID
	PhotoId   uuid.UUID
	Message   string
	CreatedAt int32
	UpdatedAt int32
	DeletedAt int32
}
