package repository

import (
	"context"

	"github.com/dihanto/gosnap/model/domain"
)

type SocialMediaRepository interface {
	PostSocialMedia(ctx context.Context, socialMedia domain.SocialMedia) (domain.SocialMedia, error)
	GetSocialMedia(ctx context.Context) ([]domain.SocialMedia, []domain.User, error)
	UpdateSocialMedia(ctx context.Context, socialMedia domain.SocialMedia) (domain.SocialMedia, error)
	DeleteSocialMedia(ctx context.Context, id int) error
}
