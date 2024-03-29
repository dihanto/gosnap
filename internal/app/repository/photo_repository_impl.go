package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/dihanto/gosnap/internal/app/helper"
	"github.com/dihanto/gosnap/model/domain"
	"github.com/google/uuid"
)

type PhotoRepositoryImpl struct {
	Database *sql.DB
}

func NewPhotoRepository(database *sql.DB) PhotoRepository {
	return &PhotoRepositoryImpl{
		Database: database,
	}
}

// PostPhoto is a method to create a new photo entry in the database.
func (repository *PhotoRepositoryImpl) PostPhoto(ctx context.Context, photo domain.Photo) (domain.Photo, error) {
	tx, err := repository.Database.Begin()
	if err != nil {
		return domain.Photo{}, err
	}
	defer helper.CommitOrRollback(tx, &err)

	photo.CreatedAt = int32(time.Now().Unix())

	query := "INSERT INTO photos( title, caption, photo_base64, user_id, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	row := tx.QueryRowContext(ctx, query, photo.Title, photo.Caption, photo.PhotoBase64, photo.UserId, photo.CreatedAt)
	err = row.Scan(&photo.Id)
	if err != nil {
		return domain.Photo{}, err
	}
	queryLike := "INSERT INTO likes (photo_id) VALUES ($1)"
	_, err = tx.ExecContext(ctx, queryLike, photo.Id)
	if err != nil {
		return domain.Photo{}, err
	}

	return photo, nil
}

// GetPhoto is a method to retrieve all photo entries and their associated users from the database.
func (repository *PhotoRepositoryImpl) GetPhoto(ctx context.Context, limit int, offset int) ([]domain.Photo, []domain.User, []domain.Like, error) {
	tx, err := repository.Database.Begin()
	if err != nil {
		return []domain.Photo{}, []domain.User{}, []domain.Like{}, err
	}
	defer helper.CommitOrRollback(tx, &err)

	query := "SELECT photos.id, photos.title, photos.caption, photos.photo_base64, photos.user_id, photos.created_at, photos.updated_at, users.username, users.email, users.profile_picture_base64, likes.like_count FROM photos JOIN users ON photos.user_id = users.id JOIN likes ON photos.id = likes.photo_id WHERE photos.deleted_at IS NULL"
	params := []interface{}{}
	if offset >= 0 {
		query += " LIMIT $1 OFFSET $2"
		params = append(params, limit, offset)
	}
	rows, err := tx.QueryContext(ctx, query, params...)
	if err != nil {
		return []domain.Photo{}, []domain.User{}, []domain.Like{}, err
	}
	defer rows.Close()

	var users []domain.User
	var photos []domain.Photo
	var likes []domain.Like
	for rows.Next() {
		photo := domain.Photo{}
		user := domain.User{}
		like := domain.Like{}
		err := rows.Scan(&photo.Id, &photo.Title, &photo.Caption, &photo.PhotoBase64, &photo.UserId, &photo.CreatedAt, &photo.UpdatedAt, &user.Username, &user.Email, &user.ProfilePicture, &like.LikeCount)
		if err != nil {
			return []domain.Photo{}, []domain.User{}, []domain.Like{}, err
		}
		user.Id = photo.UserId
		like.PhotoId = photo.Id
		users = append(users, user)
		photos = append(photos, photo)
		likes = append(likes, like)
	}
	return photos, users, likes, nil
}

// UpdatePhoto is a method to update a photo entry in the database.
func (repository *PhotoRepositoryImpl) UpdatePhoto(ctx context.Context, photo domain.Photo) (domain.Photo, error) {
	tx, err := repository.Database.Begin()
	if err != nil {
		return domain.Photo{}, err
	}
	defer helper.CommitOrRollback(tx, &err)

	photo.UpdatedAt = int32(time.Now().Unix())

	var userIdResponse uuid.UUID
	query := "UPDATE photos SET caption=$1, updated_at=$2 WHERE id=$3 RETURNING user_id"
	row := tx.QueryRowContext(ctx, query, photo.Caption, photo.UpdatedAt, photo.Id)
	err = row.Scan(&userIdResponse)
	if err != nil {
		return domain.Photo{}, err
	}
	if photo.UserId != userIdResponse {
		errr := errors.New("unauthorized")
		return domain.Photo{}, errr
	}

	return photo, nil
}

// DeletePhoto is a method to "soft delete" a photo entry by setting the deleted_at field in the database.
func (repository *PhotoRepositoryImpl) DeletePhoto(ctx context.Context, id int) error {
	tx, err := repository.Database.Begin()
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx, &err)

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

func (repository *PhotoRepositoryImpl) GetPhotoById(ctx context.Context, photoId int) (domain.Photo, domain.User, domain.Like, error) {
	tx, err := repository.Database.Begin()
	if err != nil {
		return domain.Photo{}, domain.User{}, domain.Like{}, err
	}
	defer helper.CommitOrRollback(tx, &err)

	photo := domain.Photo{}
	photoQuery := "SELECT caption, photo_base64, user_id FROM photos WHERE id=$1"
	err = tx.QueryRowContext(ctx, photoQuery, photoId).Scan(&photo.Caption, &photo.PhotoBase64, &photo.UserId)
	if err != nil {
		return domain.Photo{}, domain.User{}, domain.Like{}, err
	}

	user := domain.User{}
	userQuery := "SELECT username, name, profile_picture_base64 FROM users WHERE id=$1"
	err = tx.QueryRowContext(ctx, userQuery, photo.UserId).Scan(&user.Username, &user.Name, &user.ProfilePicture)
	if err != nil {
		return domain.Photo{}, domain.User{}, domain.Like{}, err
	}

	like := domain.Like{}
	likeQuery := "SELECT like_count from likes WHERE photo_id=$1"
	err = tx.QueryRowContext(ctx, likeQuery, photoId).Scan(&like.LikeCount)
	if err != nil {
		return domain.Photo{}, domain.User{}, domain.Like{}, err
	}

	return photo, user, like, nil

}
