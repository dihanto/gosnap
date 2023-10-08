package usecase

import (
	"context"

	"github.com/dihanto/gosnap/model/web/request"
	"github.com/dihanto/gosnap/model/web/response"
	"github.com/google/uuid"
)

type UserUsecase interface {
	UserRegister(ctx context.Context, request request.UserRegister) (response.UserRegister, error)
	UserLogin(ctx context.Context, username, password string) (bool, uuid.UUID, error)
	UserUpdate(ctx context.Context, request request.UserUpdate) (response.UserUpdate, error)
	UserDelete(ctx context.Context, id uuid.UUID) error
	FindUser(ctx context.Context, id uuid.UUID) (response.FindUser, error)
}
