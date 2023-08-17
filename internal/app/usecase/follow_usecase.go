package usecase

import (
	"context"

	"github.com/dihanto/gosnap/model/domain"
	"github.com/google/uuid"
)

type FollowUsecase interface {
	FollowUser(ctx context.Context, followerId uuid.UUID, username string) (user domain.User, err error)
	UnFollowUser(ctx context.Context, id int) (err error)
}
