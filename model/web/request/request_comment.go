package request

import "github.com/google/uuid"

type Comment struct {
	Id      uuid.UUID `json:"id"`
	UserId  uuid.UUID `json:"userId"`
	PhotoId uuid.UUID `json:"photoId"`
	Message string    `validate:"required" json:"message"`
}
