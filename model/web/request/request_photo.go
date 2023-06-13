package request

import "github.com/google/uuid"

type Photo struct {
	Id       uuid.UUID `json:"id"`
	Title    string    `validate:"required" json:"title"`
	Caption  string    `json:"caption"`
	PhotoUrl string    `validate:"required" json:"photoUrl"`
	UserId   uuid.UUID `json:"userId"`
}
