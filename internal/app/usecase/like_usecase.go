package usecase

import (
	"context"

	"github.com/dihanto/gosnap/model/web/request"
	"github.com/dihanto/gosnap/model/web/response"
)

type LikeUsecase interface {
	LikePhoto(ctx context.Context, request request.Like) (response.Like, error)
	UnlikePhoto(ctx context.Context, request request.Like) (response.Unlike, error)
	IsLikePhoto(ctx context.Context, request request.Like) (bool, error)
}
