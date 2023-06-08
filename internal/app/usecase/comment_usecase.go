package usecase

import (
	"context"

	"github.com/dihanto/gosnap/model/web"
)

type CommentUsecase interface {
	PostComment(ctx context.Context, request web.PostComment) (web.PostComment, error)
	GetComment(ctx context.Context) ([]web.GetComment, error)
	UpdateComment(ctx context.Context, request web.UpdateComment) (web.UpdateComment, error)
	DeleteComment(ctx context.Context, id int) error
}
