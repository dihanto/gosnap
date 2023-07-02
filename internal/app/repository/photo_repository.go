package repository

import (
	"context"
	"database/sql"

	"github.com/dihanto/gosnap/model/domain"
	"github.com/google/uuid"
)

type PhotoRepository interface {
	PostPhoto(ctx context.Context, tx *sql.Tx, photo domain.Photo) (domain.Photo, error)
	GetPhoto(ctx context.Context, tx *sql.Tx) ([]domain.Photo, []domain.User, error)
	UpdatePhoto(ctx context.Context, tx *sql.Tx, photo domain.Photo) (domain.Photo, error)
	DeletePhoto(ctx context.Context, tx *sql.Tx, id int) error
	LikePhoto(ctx context.Context, tx *sql.Tx, id int, userId uuid.UUID) (domain.Photo, error)
	UnLikePhoto(ctx context.Context, tx *sql.Tx, id int, userId uuid.UUID) (domain.Photo, error)
}
