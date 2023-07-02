package domain

import "github.com/google/uuid"

type LikePhoto struct {
	PhotoId int
	UserId  uuid.UUID
}
