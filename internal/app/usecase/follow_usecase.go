package usecase

import (
	"context"

	"github.com/dihanto/gosnap/model/web/request"
	"github.com/dihanto/gosnap/model/web/response"
)

type FollowUsecase interface {
	FollowUser(ctx context.Context, request request.Follow) (response.Follow, error)
	UnFollowUser(ctx context.Context, request request.Follow) (response.Follow, error)
	GetFollower(ctx context.Context, request request.Follow) (followers response.GetFollower, err error)
	GetFollowing(ctx context.Context, request request.Follow) (follows []response.GetFollowing, err error)
}
