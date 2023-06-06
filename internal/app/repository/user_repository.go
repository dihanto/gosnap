package repository

import (
	"context"
	"database/sql"

	"github.com/dihanto/gosnap/model/domain"
)

type UserRepository interface {
	UserRegister(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error)
	UserLogin(ctx context.Context, tx *sql.Tx, username string, password string) (bool, int, error)
	UserUpdate(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error)
	UserDelete(ctx context.Context, tx *sql.Tx, id int) error
}
