package response

import (
	"time"

	"github.com/google/uuid"
)

type PostComment struct {
	Id        uuid.UUID `json:"id"`
	UserId    uuid.UUID `json:"userId"`
	PhotoId   uuid.UUID `json:"photoId"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"createdAt"`
}
type UpdateComment struct {
	Id        uuid.UUID `json:"id"`
	UserId    uuid.UUID `json:"userId"`
	PhotoId   uuid.UUID `json:"photoId"`
	Message   string    `json:"message"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type GetComment struct {
	Id        uuid.UUID `json:"id"`
	Message   string    `json:"message"`
	UserId    uuid.UUID `json:"userId"`
	PhotoId   uuid.UUID `json:"photoId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	User      UserComment
	Photo     PhotoComment
}

type UserComment struct {
	Id       uuid.UUID `json:"id"`
	Email    string    `json:"emai"`
	Username string    `json:"username"`
}
type PhotoComment struct {
	Id       uuid.UUID `json:"id"`
	Title    string    `json:"title"`
	Caption  string    `json:"caption"`
	PhotoUrl string    `json:"photoUrl"`
	UserId   uuid.UUID `json:"UserId"`
}
