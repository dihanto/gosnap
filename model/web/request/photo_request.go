package request

import "github.com/google/uuid"

type Photo struct {
	Id          int       `json:"id"`
	Title       string    `validate:"required" json:"title"`
	Caption     string    `json:"caption"`
	PhotoBase64 string    `validate:"required" json:"photoBase64"`
	UserId      uuid.UUID `json:"userId"`
}
