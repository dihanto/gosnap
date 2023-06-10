package web

import "time"

type PostSocialMedia struct {
	Id             int       `json:"id"`
	Name           string    `validate:"required" json:"name"`
	SocialMediaUrl string    `validate:"required" json:"socialMediaUrl"`
	UserId         int       `json:"userId"`
	CreatedAt      time.Time `json:"createdAt"`
}

type GetSocialMedia struct {
	Id             int       `json:"id"`
	Name           string    `validate:"required" json:"name"`
	SocialMediaUrl string    `json:"socialMediaUrl"`
	UserId         int       `json:"userId"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	User           UserSocialMedia
}
type UpdateSocialMedia struct {
	Id             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `validate:"required" json:"socialMediaUrl"`
	UserId         int       `json:"userId"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
type UserSocialMedia struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}
