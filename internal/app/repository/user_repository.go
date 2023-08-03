package repository

import (
	"context"

	"github.com/dihanto/gosnap/model/domain"
	"github.com/google/uuid"
)

type UserRepository interface {
	UserRegister(ctx context.Context, user domain.User) (domain.User, error)
	UserLogin(ctx context.Context, username string, password string) (bool, uuid.UUID, error)
	UserUpdate(ctx context.Context, user domain.User) (domain.User, error)
	UserDelete(ctx context.Context, id uuid.UUID) error
}
