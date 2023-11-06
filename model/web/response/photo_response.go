package response

import (
	"time"

	"github.com/google/uuid"
)

type PostPhoto struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Caption     string    `json:"caption"`
	PhotoBase64 string    `json:"photoBase64"`
	UserId      uuid.UUID `json:"userId"`
	CreatedAt   time.Time `json:"createdAt"`
}

type UpdatePhoto struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Caption     string    `json:"caption"`
	PhotoBase64 string    `json:"photoBase64"`
	UserId      uuid.UUID `json:"userId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type GetPhoto struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Caption     string    `json:"caption"`
	PhotoBase64 string    `json:"photoBase64"`
	UserId      uuid.UUID `json:"userId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	User        User      `json:"user"`
	Likes       Likes     `json:"like"`
}

type Likes struct {
	LikeCount int `json:"likeCount"`
}

type User struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	ProfilePicture string `json:"profilePicture"`
}
