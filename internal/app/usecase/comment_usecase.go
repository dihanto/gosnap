package usecase

import (
	"context"

	"github.com/dihanto/gosnap/model/web/request"
	"github.com/dihanto/gosnap/model/web/response"
)

type CommentUsecase interface {
	PostComment(ctx context.Context, request request.Comment) (response.PostComment, error)
	GetComment(ctx context.Context) ([]response.GetComment, error)
	UpdateComment(ctx context.Context, request request.Comment) (response.UpdateComment, error)
	DeleteComment(ctx context.Context, id int) error
}
