package request

import "github.com/google/uuid"

type Photo struct {
	Id          int       `validate:"required" json:"id"`
	Title       string    `json:"title"`
	Caption     string    `json:"caption"`
	PhotoBase64 string    `json:"photoBase64"`
	UserId      uuid.UUID `json:"userId"`
}
