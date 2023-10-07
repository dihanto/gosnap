package repository

import (
	"context"
	"database/sql"

	"github.com/dihanto/gosnap/internal/app/helper"
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

func (repository *FollowRepositoryImpl) FollowUser(ctx context.Context, followerId uuid.UUID, username string) (err error) {
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

	queryFollow := "INSERT INTO follower_details(username, follower_id) VALUES ($1, $2)"
	_, err = tx.ExecContext(ctx, queryFollow, username, followerId)
	if err != nil {
		return
	}

	return
}

func (repository *FollowRepositoryImpl) UnFollowUser(ctx context.Context, followerId uuid.UUID, username string) error {
	tx, err := repository.Database.Begin()
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx, &err)

	query := "DELETE FROM follower_details WHERE follower_id=$1"
	_, err = tx.ExecContext(ctx, query, followerId)
	if err != nil {
		return err
	}

	queryUser := "UPDATE users SET followers=followers-1 WHERE username=$1"
	_, err = tx.ExecContext(ctx, queryUser, username)
	if err != nil {
		return err
	}

	return nil

}
