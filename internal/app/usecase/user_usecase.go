package usecase

import (
	"context"

	"github.com/dihanto/gosnap/model/web"
)

type UserUsecase interface {
	UserRegister(ctx context.Context, request web.UserRegister) (web.UserRegister, error)
	UserLogin(ctx context.Context, username, password string) (bool, int, error)
	UserUpdate(ctx context.Context, request web.UserUpdate) (web.UserUpdate, error)
	UserDelete(ctx context.Context, id int) error
}
