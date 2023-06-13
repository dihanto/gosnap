package repository

import (
	"context"
	"database/sql"

	"github.com/dihanto/gosnap/model/domain"
	"github.com/google/uuid"
)

type CommentRepository interface {
	PostComment(ctx context.Context, tx *sql.Tx, comment domain.Comment) (domain.Comment, error)
	GetComment(ctx context.Context, tx *sql.Tx) ([]domain.Comment, []domain.User, []domain.Photo, error)
	UpdateComment(ctx context.Context, tx *sql.Tx, comment domain.Comment) (domain.Comment, error)
	DeleteComment(ctx context.Context, tx *sql.Tx, id uuid.UUID) error
}
