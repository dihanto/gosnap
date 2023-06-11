package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/dihanto/gosnap/model/domain"
)

type PhotoRepositoryImpl struct {
}

func NewPhotoRepository() PhotoRepository {
	return &PhotoRepositoryImpl{}
}

func (repository *PhotoRepositoryImpl) PostPhoto(ctx context.Context, tx *sql.Tx, photo domain.Photo) (domain.Photo, error) {
	t := time.Now()
	photo.CreatedAt = int32(t.Unix())

	query := "insert into photos(title, caption, photo_url,user_id, created_at) values ($1, $2, $3, $4, $5) returning id"
	row := tx.QueryRowContext(ctx, query, photo.Title, photo.Caption, photo.PhotoUrl, photo.UserId, photo.CreatedAt)

	var id int
	err := row.Scan(&id)
	if err != nil {
		return domain.Photo{}, err
	}
	photo.Id = id

	return photo, nil

}

func (repository *PhotoRepositoryImpl) GetPhoto(ctx context.Context, tx *sql.Tx) ([]domain.Photo, domain.User, error) {
	query := "select photos.id, photos.title, photos.caption, photos.photo_url, photos.user_id, photos.created_at, photos.updated_at, users.username, users.email from photos	join users on photos.user_id = users.id where photos.deleted_at is null;"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []domain.Photo{}, domain.User{}, err
	}
	defer rows.Close()
	user := domain.User{}
	var photos []domain.Photo
	for rows.Next() {
		photo := domain.Photo{}
		err := rows.Scan(&photo.Id, &photo.Title, &photo.Caption, &photo.PhotoUrl, &photo.UserId, &photo.CreatedAt, &photo.UpdatedAt, &user.Username, &user.Email)
		if err != nil {
			return []domain.Photo{}, domain.User{}, err
		}
		photos = append(photos, photo)
	}

	return photos, user, nil
}

func (repository *PhotoRepositoryImpl) UpdatePhoto(ctx context.Context, tx *sql.Tx, photo domain.Photo) (domain.Photo, error) {
	t := time.Now()
	photo.UpdatedAt = int32(t.Unix())

	query := "update photos set title=$1, caption=$2, photo_url=$3, user_id=$4, updated_at=$5 where id=$6 returning created_at"
	row := tx.QueryRowContext(ctx, query, photo.Title, photo.Caption, photo.PhotoUrl, photo.UserId, photo.UpdatedAt, photo.Id)

	err := row.Scan(&photo.CreatedAt)
	if err != nil {
		return domain.Photo{}, err
	}

	return photo, nil

}

func (repository *PhotoRepositoryImpl) DeletePhoto(ctx context.Context, tx *sql.Tx, id int) error {
	t := time.Now()
	deleteTime := t.Unix()

	query := "update photos set deleted_at=$1 where id=$2"
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
