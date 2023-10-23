package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/dihanto/gosnap/internal/app/helper"
	"github.com/dihanto/gosnap/model/domain"
)

type CommentRepositoryImpl struct {
	Database *sql.DB
}

func NewCommentRepository(database *sql.DB) CommentRepository {
	return &CommentRepositoryImpl{
		Database: database,
	}
}

// PostComment is a method to create a new comment entry in the database.
func (repository *CommentRepositoryImpl) PostComment(ctx context.Context, comment domain.Comment) (domain.Comment, error) {
	tx, err := repository.Database.Begin()
	if err != nil {
		return domain.Comment{}, nil
	}
	defer helper.CommitOrRollback(tx, &err)

	comment.CreatedAt = int32(time.Now().Unix())

	query := "INSERT INTO comments (message, photo_id, user_id, created_at) VALUES ($1, $2, $3, $4) RETURNING id"
	row := tx.QueryRowContext(ctx, query, comment.Message, comment.PhotoId, comment.UserId, comment.CreatedAt)
	err = row.Scan(&comment.Id)
	if err != nil {
		return domain.Comment{}, err
	}

	return comment, nil
}

// GetComment is a method to retrieve all comment entries and their associated users and photos from the database.
func (repository *CommentRepositoryImpl) GetComment(ctx context.Context) ([]domain.Comment, []domain.User, []domain.Photo, error) {
	tx, err := repository.Database.Begin()
	if err != nil {
		return []domain.Comment{}, []domain.User{}, []domain.Photo{}, nil
	}
	defer helper.CommitOrRollback(tx, &err)

	query := "SELECT comments.id, comments.message, comments.photo_id, comments.user_id, comments.created_at, comments.updated_at, users.id, users.email, users.username, photos.id, photos.title, photos.caption, photos.photo_base64, photos.user_id FROM comments JOIN photos ON comments.photo_id = photos.id JOIN users ON comments.user_id = users.id WHERE comments.deleted_at IS NULL;"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []domain.Comment{}, []domain.User{}, []domain.Photo{}, err
	}
	defer rows.Close()

	comments := []domain.Comment{}
	var users []domain.User
	var photos []domain.Photo

	for rows.Next() {
		var comment domain.Comment
		var user domain.User
		var photo domain.Photo
		err = rows.Scan(&comment.Id, &comment.Message, &comment.PhotoId, &comment.UserId, &comment.CreatedAt, &comment.UpdatedAt, &user.Id, &user.Email, &user.Username, &photo.Id, &photo.Title, &photo.Caption, &photo.PhotoBase64, &photo.UserId)
		if err != nil {
			return []domain.Comment{}, []domain.User{}, []domain.Photo{}, err
		}
		user.Id = comment.UserId
		photo.Id = comment.PhotoId
		users = append(users, user)
		photos = append(photos, photo)
		comments = append(comments, comment)
	}

	return comments, users, photos, nil
}

// UpdateComment is a method to update a comment entry in the database.
func (repository *CommentRepositoryImpl) UpdateComment(ctx context.Context, comment domain.Comment) (domain.Comment, error) {
	tx, err := repository.Database.Begin()
	if err != nil {
		return domain.Comment{}, nil
	}
	defer helper.CommitOrRollback(tx, &err)

	comment.UpdatedAt = int32(time.Now().Unix())

	query := "UPDATE comments SET message=$1, updated_at=$2, user_id=$3 WHERE id=$4 RETURNING photo_id"
	row := tx.QueryRowContext(ctx, query, comment.Message, comment.UpdatedAt, comment.UserId, comment.Id)

	err = row.Scan(&comment.PhotoId)
	if err != nil {
		return domain.Comment{}, err
	}

	return comment, nil
}

// DeleteComment is a method to "soft delete" a comment entry by setting the deleted_at field in the database.
func (repository *CommentRepositoryImpl) DeleteComment(ctx context.Context, id int) error {
	tx, err := repository.Database.Begin()
	if err != nil {
		return nil
	}
	defer helper.CommitOrRollback(tx, &err)

	deleteTime := int32(time.Now().Unix())

	query := "UPDATE comments SET deleted_at = $1 WHERE id = $2"
	result, err := tx.ExecContext(ctx, query, deleteTime, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("comment not found")
	}

	return nil
}
