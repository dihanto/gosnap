package repository

import (
	"context"

	"github.com/dihanto/gosnap/model/domain"
)

type PhotoRepository interface {
	PostPhoto(ctx context.Context, photo domain.Photo) (domain.Photo, error)
	GetPhoto(ctx context.Context, limit int, offset int) ([]domain.Photo, []domain.User, []domain.Like, error)
	UpdatePhoto(ctx context.Context, photo domain.Photo) (domain.Photo, error)
	DeletePhoto(ctx context.Context, id int) error
}
