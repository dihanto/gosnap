package domain

import "github.com/google/uuid"

type SocialMedia struct {
	Id             uuid.UUID
	Name           string
	SocialMediaUrl string
	UserId         uuid.UUID
	CreatedAt      int32
	UpdatedAt      int32
	DeletedAt      int32
}
