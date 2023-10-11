package request

import "github.com/google/uuid"

type Like struct {
	PhotoId int       `json:"photoId" validate:"required"`
	UserId  uuid.UUID `json:"userId" validate:"required"`
}
