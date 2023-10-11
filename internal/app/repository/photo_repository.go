package repository

import (
	"context"

	"github.com/dihanto/gosnap/model/domain"
)

type PhotoRepository interface {
	PostPhoto(ctx context.Context, photo domain.Photo) (domain.Photo, error)
	GetPhoto(ctx context.Context) ([]domain.Photo, []domain.User, error)
	UpdatePhoto(ctx context.Context, photo domain.Photo) (domain.Photo, error)
	DeletePhoto(ctx context.Context, id int) error
}
