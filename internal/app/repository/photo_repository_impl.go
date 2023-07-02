package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/dihanto/gosnap/model/domain"
	"github.com/google/uuid"
)

type PhotoRepositoryImpl struct {
}

func NewPhotoRepository() PhotoRepository {
	return &PhotoRepositoryImpl{}
}

func (repository *PhotoRepositoryImpl) PostPhoto(ctx context.Context, tx *sql.Tx, photo domain.Photo) (domain.Photo, error) {
	photo.CreatedAt = int32(time.Now().Unix())

	query := "INSERT INTO photos( title, caption, photo_url, user_id, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	row := tx.QueryRowContext(ctx, query, photo.Title, photo.Caption, photo.PhotoUrl, photo.UserId, photo.CreatedAt)
	err := row.Scan(&photo.Id)
	if err != nil {
		return domain.Photo{}, err
	}

	return photo, nil

}

func (repository *PhotoRepositoryImpl) GetPhoto(ctx context.Context, tx *sql.Tx) ([]domain.Photo, []domain.User, error) {

	query := "SELECT photos.id, photos.title, photos.caption, photos.photo_url, photos.user_id, photos.created_at, photos.updated_at, users.username, users.email FROM photos JOIN users ON photos.user_id = users.id WHERE photos.deleted_at IS NULL;"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []domain.Photo{}, []domain.User{}, err
	}
	defer rows.Close()
	var users []domain.User
	var photos []domain.Photo
	for rows.Next() {
		photo := domain.Photo{}
		user := domain.User{}
		err := rows.Scan(&photo.Id, &photo.Title, &photo.Caption, &photo.PhotoUrl, &photo.UserId, &photo.CreatedAt, &photo.UpdatedAt, &user.Username, &user.Email)
		if err != nil {
			return []domain.Photo{}, []domain.User{}, err
		}
		user.Id = photo.UserId
		users = append(users, user)
		photos = append(photos, photo)
	}

	return photos, users, nil
}

func (repository *PhotoRepositoryImpl) UpdatePhoto(ctx context.Context, tx *sql.Tx, photo domain.Photo) (domain.Photo, error) {
	photo.UpdatedAt = int32(time.Now().Unix())

	query := "UPDATE photos SET title=$1, caption=$2, photo_url=$3, user_id=$4, updated_at=$5 WHERE id=$6 RETURNING created_at"
	row := tx.QueryRowContext(ctx, query, photo.Title, photo.Caption, photo.PhotoUrl, photo.UserId, photo.UpdatedAt, photo.Id)

	err := row.Scan(&photo.CreatedAt)
	if err != nil {
		return domain.Photo{}, err
	}

	return photo, nil

}

func (repository *PhotoRepositoryImpl) DeletePhoto(ctx context.Context, tx *sql.Tx, id int) error {
	deleteTime := int32(time.Now().Unix())

	query := "UPDATE photos SET deleted_at=$1 WHERE id=$2"
	result, err := tx.ExecContext(ctx, query, deleteTime, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("photo media not found")
	}

	return nil
}

func (repository *PhotoRepositoryImpl) LikePhoto(ctx context.Context, tx *sql.Tx, id int, userId uuid.UUID) (domain.Photo, error) {
	query := "UPDATE photos SET likes=likes+1 WHERE id=$1"
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return domain.Photo{}, err
	}
	queryLike := "INSERT INTO like_details (photo_id, user_id) VALUES ($1, $2)"
	_, err = tx.ExecContext(ctx, queryLike, id, userId)
	if err != nil {
		return domain.Photo{}, err
	}
	queryResult := "SELECT title, photo_url, likes FROM photos WHERE id=$1"
	rows, err := tx.QueryContext(ctx, queryResult, id)
	if err != nil {
		return domain.Photo{}, err
	}
	defer rows.Close()

	var photo domain.Photo
	if rows.Next() {
		err = rows.Scan(&photo.Title, &photo.PhotoUrl, &photo.Likes)
		if err != nil {
			return domain.Photo{}, err
		}
	}
	photo.Id = id

	return photo, err

}

func (repository *PhotoRepositoryImpl) UnLikePhoto(ctx context.Context, tx *sql.Tx, id int, userId uuid.UUID) (domain.Photo, error) {
	panic("not implemented") // TODO: Implement
}
