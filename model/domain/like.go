package domain

import "github.com/google/uuid"

type Like struct {
	Id        int
	LikeCount int
	PhotoId   int
	UserId    uuid.UUID
	CreatedAt int32
}
