package usecase

import (
	"context"

	"github.com/dihanto/gosnap/model/web/request"
	"github.com/dihanto/gosnap/model/web/response"
)

type SocialMediaUsecase interface {
	PostSocialMedia(ctx context.Context, request request.SocialMedia) (response.PostSocialMedia, error)
	GetSocialMedia(ctx context.Context) ([]response.GetSocialMedia, error)
	UpdateSocialMedia(ctx context.Context, request request.SocialMedia) (response.UpdateSocialMedia, error)
	DeleteSocialMedia(ctx context.Context, id int) error
}
