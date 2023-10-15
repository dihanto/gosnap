package repository

import (
	"context"

	"github.com/dihanto/gosnap/model/domain"
)

type FollowRepository interface {
	FollowUser(ctx context.Context, follow domain.Follow) (domain.Follow, error)
	UnFollowUser(ctx context.Context, follow domain.Follow) (domain.Follow, error)
	GetFollower(ctx context.Context, request domain.Follow) (followers []domain.User, err error)
	GetFollowing(ctx context.Context, request domain.Follow) (follows []domain.User, err error)
}
