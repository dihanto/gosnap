package request

import "github.com/google/uuid"

type Comment struct {
	Id      int       `json:"id"`
	UserId  uuid.UUID `json:"userId"`
	PhotoId int       `json:"photoId"`
	Message string    `validate:"required" json:"message"`
}
