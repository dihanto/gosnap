package repository

import (
	"context"
	"database/sql"

	"github.com/dihanto/gosnap/internal/app/helper"
	"github.com/dihanto/gosnap/model/domain"
	"github.com/google/uuid"
)

type FollowRepositoryImpl struct {
	Database *sql.DB
}

func NewFollowRepositoryImpl(database *sql.DB) FollowRepository {
	return &FollowRepositoryImpl{
		Database: database,
	}
}

func (repository *FollowRepositoryImpl) FollowUser(ctx context.Context, followerId uuid.UUID, username string) (user domain.User, err error) {
	tx, err := repository.Database.Begin()
	if err != nil {
		return
	}
	defer helper.CommitOrRollback(tx, &err)

	query := "UPDATE users SET followers = followers + 1 WHERE username = $1"
	_, err = tx.ExecContext(ctx, query, username)
	if err != nil {
		return
	}

	return
}

func (repository *FollowRepositoryImpl) UnFollowUser(ctx context.Context, id int) error {
	panic("not implemented") // TODO: Implement
}
