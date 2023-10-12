package response

import (
	"time"

	"github.com/google/uuid"
)

type Like struct {
	PhotoId   int       `json:"photoId"`
	UserId    uuid.UUID `json:"userId"`
	LikeCount int       `json:"likeCount"`
	LikedAt   time.Time `json:"LikedAt"`
}

type Unlike struct {
	PhotoId   int `json:"photoId"`
	LikeCount int `json:"likeCount"`
}
