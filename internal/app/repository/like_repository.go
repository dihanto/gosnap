package repository

import (
	"context"

	"github.com/dihanto/gosnap/model/domain"
	"github.com/google/uuid"
)

type LikeRepository interface {
	LikePhoto(ctx context.Context, like domain.Like) (domain.Like, error)
	UnlikePhoto(ctx context.Context, like domain.Like) (domain.Like, error)
	IsLikePhoto(ctx context.Context, photoId int) ([]uuid.UUID, error)
}
