package usecase

import (
	"context"

	"github.com/dihanto/gosnap/model/web/request"
	"github.com/dihanto/gosnap/model/web/response"
	"github.com/google/uuid"
)

type PhotoUsecase interface {
	PostPhoto(ctx context.Context, request request.Photo) (response.PostPhoto, error)
	GetPhoto(ctx context.Context) ([]response.GetPhoto, error)
	UpdatePhoto(ctx context.Context, request request.Photo) (response.UpdatePhoto, error)
	DeletePhoto(ctx context.Context, id uuid.UUID) error
}
