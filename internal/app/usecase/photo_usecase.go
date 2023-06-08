package usecase

import (
	"context"

	"github.com/dihanto/gosnap/model/web"
)

type PhotoUsecase interface {
	PostPhoto(ctx context.Context, request web.Photo) (web.Photo, error)
	GetPhoto(ctx context.Context) ([]web.GetPhoto, error)
	UpdatePhoto(ctx context.Context, request web.Photo) (web.UpdatePhoto, error)
	DeletePhoto(ctx context.Context, id int) error
}
