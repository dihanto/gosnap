package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/dihanto/gosnap/model/domain"
)

type PhotoRepositoryImpl struct {
}

func (repository *PhotoRepositoryImpl) CreatePhoto(ctx context.Context, tx *sql.Tx, photo domain.Photo) (domain.Photo, error) {
	t := time.Now()
	photo.CreatedAt = int32(t.Unix())

	query := "insert into photos(title, caption, photo_url, created_at) values ($1, $2, $3)"
	tx.QueryRowContext(ctx, query, photo.Title, photo.Caption, photo.PhotoUrl, photo.CreatedAt)

	return photo, nil

}

func (repository *PhotoRepositoryImpl) GetPhoto(ctx context.Context, tx *sql.Tx) ([]domain.Photo, error) {
	query := "select id, title, caption, photo_url, user_id, created_at, updated_at from photos"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var photos []domain.Photo
	for rows.Next() {
		photo := domain.Photo{}
		err := rows.Scan(&photo.Id, &photo.Title, &photo.Caption, &photo.User_id, &photo.CreatedAt, &photo.UpdatedAt)
		if err != nil {
			panic(err)
		}
		photos = append(photos, photo)
	}

	return photos, nil
}

func (repository *PhotoRepositoryImpl) UpdatePhoto(ctx context.Context, tx *sql.Tx, photo domain.Photo) (domain.Photo, error) {
	t := time.Now()
	photo.UpdatedAt = int32(t.Unix())

	query := "update from photos set title=$1, caption=$2, photo_url=$3, updated_at=$4 where id=$5"
	row := tx.QueryRowContext(ctx, query, photo.Title, photo.Caption, photo.PhotoUrl, photo.UpdatedAt, photo.Id)

	row.Scan(&photo.User_id)

	return photo, nil

}

func (repository *PhotoRepositoryImpl) DeletePhoto(ctx context.Context, tx *sql.Tx, id int) error {
	t := time.Now()
	deleteTime := t.Unix()

	query := "update from photos set deleted_at=$1 where id=$2"
	_, err := tx.ExecContext(ctx, query, deleteTime, id)
	if err != nil {
		panic(err)
	}

	return nil
}
