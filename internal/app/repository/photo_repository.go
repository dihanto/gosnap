package repository

import (
	"context"
	"database/sql"

	"github.com/dihanto/gosnap/model/domain"
)

type PhotoRepository interface {
	CreatePhoto(ctx context.Context, tx *sql.Tx, photo domain.Photo) (domain.Photo, error)
	GetPhoto(ctx context.Context, tx *sql.Tx) ([]domain.Photo, error)
	UpdatePhoto(ctx context.Context, tx *sql.Tx, photo domain.Photo) (domain.Photo, error)
	DeletePhoto(ctx context.Context, tx *sql.Tx, id int) error
}
