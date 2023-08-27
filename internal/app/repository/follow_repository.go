package repository

import (
	"context"

	"github.com/google/uuid"
)

type FollowRepository interface {
	FollowUser(ctx context.Context, followerId uuid.UUID, username string) (err error)
	UnFollowUser(ctx context.Context, id int) error
}
