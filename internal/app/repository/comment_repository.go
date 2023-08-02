package repository

import (
	"context"

	"github.com/dihanto/gosnap/model/domain"
)

type CommentRepository interface {
	PostComment(ctx context.Context, comment domain.Comment) (domain.Comment, error)
	GetComment(ctx context.Context) ([]domain.Comment, []domain.User, []domain.Photo, error)
	UpdateComment(ctx context.Context, comment domain.Comment) (domain.Comment, error)
	DeleteComment(ctx context.Context, id int) error
}
