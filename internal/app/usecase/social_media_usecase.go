package usecase

import (
	"context"

	"github.com/dihanto/gosnap/model/web"
)

type SocialMediaUsecase interface {
	PostSocialMedia(ctx context.Context, request web.PostSocialMedia) (web.PostSocialMedia, error)
	GetSocialMedia(ctx context.Context) ([]web.GetSocialMedia, error)
	UpdateSocialMedia(ctx context.Context, request web.UpdateSocialMedia) (web.UpdateSocialMedia, error)
	DeleteSocialMedia(ctx context.Context, id int) error
}
