package request

import "github.com/google/uuid"

type Photo struct {
	Id       int       `json:"id"`
	Title    string    `validate:"required" json:"title"`
	Caption  string    `json:"caption"`
	PhotoUrl string    `validate:"required" json:"photoUrl"`
	UserId   uuid.UUID `json:"userId"`
}
