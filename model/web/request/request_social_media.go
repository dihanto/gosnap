package request

import "github.com/google/uuid"

type SocialMedia struct {
	Id             uuid.UUID `json:"id"`
	Name           string    `validate:"required" json:"name"`
	SocialMediaUrl string    `validate:"required" json:"socialMediaUrl"`
	UserId         uuid.UUID `json:"userId"`
}
