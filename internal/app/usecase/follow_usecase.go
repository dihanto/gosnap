package usecase

import (
	"context"

	"github.com/google/uuid"
)

type FollowUsecase interface {
	FollowUser(ctx context.Context, followerId uuid.UUID, username string) (err error)
	UnFollowUser(ctx context.Context, id int) (err error)
}
