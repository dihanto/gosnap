package repository

import (
	"context"

	"github.com/dihanto/gosnap/model/domain"
)

type FollowRepository interface {
	FollowUser(ctx context.Context, follow domain.Follow) (domain.Follow, error)
	UnFollowUser(ctx context.Context, follow domain.Follow) (domain.Follow, error)
}
