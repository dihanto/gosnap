package domain

import "github.com/google/uuid"

type SocialMedia struct {
	Id             int
	Name           string
	SocialMediaUrl string
	UserId         uuid.UUID
	CreatedAt      int32
	UpdatedAt      int32
	DeletedAt      int32
}
