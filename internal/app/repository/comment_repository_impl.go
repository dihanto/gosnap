package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/dihanto/gosnap/model/domain"
)

type CommentRepositoryImpl struct {
}

func NewCommentRepository() CommentRepository {
	return &CommentRepositoryImpl{}
}
func (repository *CommentRepositoryImpl) PostComment(ctx context.Context, tx *sql.Tx, comment domain.Comment) (domain.Comment, error) {
	t := time.Now()
	comment.CreatedAt = int32(t.Unix())

	query := "insert into comments (message, photo_id, user_id, created_at) values($1, $2, $3, $4) returning id"
	row := tx.QueryRowContext(ctx, query, comment.Message, comment.PhotoId, comment.UserId, comment.CreatedAt)

	err := row.Scan(&comment.Id)
	if err != nil {
		panic(err)
	}

	return comment, nil
}

func (repository *CommentRepositoryImpl) GetComment(ctx context.Context, tx *sql.Tx) ([]domain.Comment, domain.User, domain.Photo, error) {
	query := "select comments.id, comments.message, comments.photo_id, comments.user_id, comments.created_at, comments.updated_at, users.id, users.email, users.username, photos.id, photos.title, photos.caption, photos.photo_url, photos.user_id	from comments	join photos on comments.photo_id = photos.id	join users on comments.user_id = users.id;"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	comments := []domain.Comment{}
	var user domain.User
	var photo domain.Photo

	for rows.Next() {
		var comment domain.Comment

		err = rows.Scan(&comment.Id, &comment.Message, &comment.PhotoId, &comment.UserId, &comment.CreatedAt, &comment.UpdatedAt, &user.Id, &user.Email, &user.Username, &photo.Id, &photo.Title, &photo.Caption, &photo.PhotoUrl, &photo.UserId)
		if err != nil {
			panic(err)
		}
		comments = append(comments, comment)
	}
	return comments, user, photo, nil
}

func (repository *CommentRepositoryImpl) UpdateComment(ctx context.Context, tx *sql.Tx, comment domain.Comment) (domain.Comment, error) {
	t := time.Now()
	comment.UpdatedAt = int32(t.Unix())

	query := "update comments set message=$1, updated_at=$2, user_id=$3 where id=$4 returning photo_id"
	row := tx.QueryRowContext(ctx, query, comment.Message, comment.UpdatedAt, comment.UserId, comment.Id)

	err := row.Scan(&comment.PhotoId)
	if err != nil {
		panic(err)
	}

	return comment, nil
}

func (repository *CommentRepositoryImpl) DeleteComment(ctx context.Context, tx *sql.Tx, id int) error {
	t := time.Now()
	deleteTime := t.Unix()

	query := "update comments set deleted_at = $1 where id = $2"
	_, err := tx.ExecContext(ctx, query, deleteTime, id)
	if err != nil {
		panic(err)
	}

	return nil
}
