package repository

import (
	"context"

	"github.com/dihanto/gosnap/model/domain"
	"github.com/google/uuid"
)

type PhotoRepository interface {
	PostPhoto(ctx context.Context, photo domain.Photo) (domain.Photo, error)
	GetPhoto(ctx context.Context) ([]domain.Photo, []domain.User, error)
	UpdatePhoto(ctx context.Context, photo domain.Photo) (domain.Photo, error)
	DeletePhoto(ctx context.Context, id int) error
	LikePhoto(ctx context.Context, id int, userId uuid.UUID) (domain.Photo, error)
	UnLikePhoto(ctx context.Context, id int, userId uuid.UUID) (domain.Photo, error)
}