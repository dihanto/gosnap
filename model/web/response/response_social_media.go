package response

import (
	"time"

	"github.com/google/uuid"
)

type PostSocialMedia struct {
	Id             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"socialMediaUrl"`
	UserId         uuid.UUID `json:"userId"`
	CreatedAt      time.Time `json:"createdAt"`
}

type GetSocialMedia struct {
	Id             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"socialMediaUrl"`
	UserId         uuid.UUID `json:"userId"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	User           UserSocialMedia
}
type UpdateSocialMedia struct {
	Id             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"socialMediaUrl"`
	UserId         uuid.UUID `json:"userId"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
type UserSocialMedia struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
}
