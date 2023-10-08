package response

import (
	"time"

	"github.com/google/uuid"
)

type PostPhoto struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photoUrl"`
	UserId    uuid.UUID `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
}

type UpdatePhoto struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photoUrl"`
	UserId    uuid.UUID `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type GetPhoto struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	Likes     int       `json:"likes"`
	PhotoUrl  string    `json:"photoUrl"`
	UserId    uuid.UUID `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	User      User      `json:"user"`
}

type LikePhoto struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	PhotoUrl string `json:"photoUrl"`
	Likes    int    `json:"likes"`
}
type UnLikePhoto struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	PhotoUrl string `json:"photoUrl"`
	Likes    int    `json:"likes"`
}

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
