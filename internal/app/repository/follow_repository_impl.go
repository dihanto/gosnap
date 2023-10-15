package repository

import (
	"context"
	"database/sql"

	"github.com/dihanto/gosnap/internal/app/helper"
	"github.com/dihanto/gosnap/model/domain"
	"github.com/lib/pq"
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

	query := "UPDATE followers SET follower_count=follower_count+1 WHERE username=$1 RETURNING id"
	err = tx.QueryRowContext(ctx, query, follow.TargetUsername).Scan(&follow.Id)
	if err != nil {
		return domain.Follow{}, err
	}

	queryFollow := "INSERT INTO follower_details(follow_id, follower_name) VALUES ($1, $2)"
	_, err = tx.ExecContext(ctx, queryFollow, follow.Id, follow.FollowerUsername)
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

func (repository *FollowRepositoryImpl) GetFollower(ctx context.Context, request domain.Follow) (followers []domain.User, err error) {
	tx, err := repository.Database.Begin()
	if err != nil {
		return
	}
	defer helper.CommitOrRollback(tx, &err)

	var followerId int
	queryGetId := "SELECT id FROM followers WHERE username=$1"
	err = tx.QueryRowContext(ctx, queryGetId, request.TargetUsername).Scan(&followerId)
	if err != nil {
		return
	}

	queryGetFollower := "SELECT follower_name FROM follower_details WHERE follow_id=$1"
	rows, err := tx.QueryContext(ctx, queryGetFollower, followerId)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var follower domain.User
		err = rows.Scan(&follower.Username)
		if err != nil {
			return
		}

		followers = append(followers, follower)
	}
	return
}

func (repository *FollowRepositoryImpl) GetFollowing(ctx context.Context, request domain.Follow) (follows []domain.User, err error) {
	tx, err := repository.Database.Begin()
	if err != nil {
		return
	}
	defer helper.CommitOrRollback(tx, &err)

	queryFollowId := "SELECT follow_id FROM follower_details WHERE follower_name=$1"
	rows, err := tx.QueryContext(ctx, queryFollowId, request.TargetUsername)
	if err != nil {
		return
	}
	defer rows.Close()

	var followIds []int
	for rows.Next() {
		var followId int
		err = rows.Scan(&followId)
		if err != nil {
			return
		}
		followIds = append(followIds, followId)
	}

	// for _, id := range followIds {
	queryFollows := "SELECT username FROM followers WHERE id= ANY ($1)"
	rowUsers, err := tx.QueryContext(ctx, queryFollows, pq.Array(followIds))
	if err != nil {
		return
	}

	for rowUsers.Next() {
		var follow domain.User
		err = rowUsers.Scan(&follow.Username)
		if err != nil {
			return
		}
		follows = append(follows, follow)
	}
	// follows = append(follows, follow)
	// }

	return
}
