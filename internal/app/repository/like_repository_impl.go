package repository

import (
	"context"
	"database/sql"

	"github.com/dihanto/gosnap/internal/app/helper"
	"github.com/dihanto/gosnap/model/domain"
	"github.com/google/uuid"
)

type LikeRepositoryImpl struct {
	Database *sql.DB
}

func NewLikeRepository(database *sql.DB) LikeRepository {
	return &LikeRepositoryImpl{
		Database: database,
	}
}
func (repository *LikeRepositoryImpl) LikePhoto(ctx context.Context, like domain.Like) (domain.Like, error) {
	tx, err := repository.Database.Begin()
	if err != nil {
		return domain.Like{}, err
	}
	defer helper.CommitOrRollback(tx, &err)

	query := "UPDATE likes SET like_count=like_count+1 WHERE photo_id=$1 RETURNING id"
	err = tx.QueryRowContext(ctx, query, like.PhotoId).Scan(&like.Id)
	if err != nil {
		return domain.Like{}, err
	}

	queryLikeDetail := "INSERT INTO like_details (like_id, user_id, liked_at) VALUES ($1, $2, $3)"
	_, err = tx.ExecContext(ctx, queryLikeDetail, like.Id, like.UserId, like.CreatedAt)
	if err != nil {
		return domain.Like{}, err
	}

	queryGetLike := "SELECT like_count FROM likes WHERE photo_id=$1"
	row, err := tx.QueryContext(ctx, queryGetLike, like.PhotoId)
	if err != nil {
		return domain.Like{}, err
	}
	defer row.Close()

	if row.Next() {
		err = row.Scan(&like.LikeCount)
		if err != nil {
			return domain.Like{}, err
		}
	}

	likeResponse := domain.Like{
		PhotoId:   like.PhotoId,
		UserId:    like.UserId,
		LikeCount: like.LikeCount,
		CreatedAt: like.CreatedAt,
	}
	return likeResponse, nil

}

func (repository *LikeRepositoryImpl) UnlikePhoto(ctx context.Context, like domain.Like) (domain.Like, error) {
	tx, err := repository.Database.Begin()
	if err != nil {
		return domain.Like{}, err
	}
	defer helper.CommitOrRollback(tx, &err)

	query := "UPDATE likes SET like_count=like_count-1 WHERE photo_id=$1"
	_, err = tx.ExecContext(ctx, query, like.PhotoId)
	if err != nil {
		return domain.Like{}, err
	}

	queryLikeDetail := "DELETE FROM like_details WHERE user_id=$1"
	_, err = tx.ExecContext(ctx, queryLikeDetail, like.UserId)
	if err != nil {
		return domain.Like{}, err
	}

	queryResult := "SELECT like_count FROM likes WHERE photo_id=$1"
	row, err := tx.QueryContext(ctx, queryResult, like.PhotoId)
	if err != nil {
		return domain.Like{}, err
	}
	defer row.Close()

	if row.Next() {
		err = row.Scan(&like.LikeCount)
		if err != nil {
			return domain.Like{}, err
		}
	}

	likeResponse := domain.Like{
		PhotoId:   like.PhotoId,
		LikeCount: like.LikeCount,
	}

	return likeResponse, nil
}

func (repository *LikeRepositoryImpl) IsLikePhoto(ctx context.Context, photoId int) ([]uuid.UUID, error) {
	tx, err := repository.Database.Begin()
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(tx, &err)

	var likeId int
	queryGetLikeId := "SELECT id FROM likes WHERE photo_id=$1"
	err = tx.QueryRowContext(ctx, queryGetLikeId, photoId).Scan(&likeId)
	if err != nil {
		return nil, err
	}

	query := "SELECT user_id FROM like_details WHERE like_id=$1"
	rows, err := tx.QueryContext(ctx, query, likeId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userIds []uuid.UUID
	for rows.Next() {
		var userId uuid.UUID
		err = rows.Scan(&userId)
		if err != nil {
			return nil, err
		}
		userIds = append(userIds, userId)
	}
	return userIds, nil
}
