package repository

import (
	"context"
	"database/sql"

	"github.com/dihanto/gosnap/internal/app/helper"
	"github.com/dihanto/gosnap/model/domain"
)

type FollowRepositoryImpl struct {
	Database *sql.DB
}

func NewFollowRepositoryImpl(database *sql.DB) FollowRepository {
	return &FollowRepositoryImpl{
		Database: database,
	}
}

func (repository *FollowRepositoryImpl) FollowUser(ctx context.Context, follow domain.Follow) (domain.Follow, error) {
	tx, err := repository.Database.Begin()
	if err != nil {
		return domain.Follow{}, err
	}
	defer helper.CommitOrRollback(tx, &err)

	query := "UPDATE followers SET follower_count=follower_count+1 WHERE username=$1"
	_, err = tx.ExecContext(ctx, query, follow.TargetUsername)
	if err != nil {
		return domain.Follow{}, err
	}

	queryFollow := "INSERT INTO follower_details(username, follower_name) VALUES ($1, $2)"
	_, err = tx.ExecContext(ctx, queryFollow, follow.TargetUsername, follow.FollowerUsername)
	if err != nil {
		return domain.Follow{}, err
	}

	queryResult := "SELECT follower_count FROM followers WHERE username=$1"
	err = tx.QueryRowContext(ctx, queryResult, follow.TargetUsername).Scan(&follow.FollowerCount)
	if err != nil {
		return domain.Follow{}, err
	}

	return follow, nil
}

func (repository *FollowRepositoryImpl) UnFollowUser(ctx context.Context, follow domain.Follow) (domain.Follow, error) {
	tx, err := repository.Database.Begin()
	if err != nil {
		return domain.Follow{}, err
	}
	defer helper.CommitOrRollback(tx, &err)

	query := "DELETE FROM follower_details WHERE follower_name=$1"
	_, err = tx.ExecContext(ctx, query, follow.FollowerUsername)
	if err != nil {
		return domain.Follow{}, err
	}

	queryUser := "UPDATE followers SET follower_count=follower_count-1 WHERE username=$1"
	_, err = tx.ExecContext(ctx, queryUser, follow.TargetUsername)
	if err != nil {
		return domain.Follow{}, err
	}

	queryResult := "SELECT follower_count FROM followers WHERE username=$1"
	err = tx.QueryRowContext(ctx, queryResult, follow.TargetUsername).Scan(&follow.FollowerCount)
	if err != nil {
		return domain.Follow{}, err
	}

	return follow, nil

}
