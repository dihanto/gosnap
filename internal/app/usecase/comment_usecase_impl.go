package usecase

import (
	"context"
	"database/sql"
	"time"

	"github.com/dihanto/gosnap/internal/app/repository"
	"github.com/dihanto/gosnap/model/domain"
	"github.com/dihanto/gosnap/model/web"
	"github.com/go-playground/validator/v10"
)

type CommentUsecaseImpl struct {
	Repository repository.CommentRepository
	DB         *sql.DB
	Validate   *validator.Validate
}

func NewCommentUsecase(repository repository.CommentRepository, db *sql.DB, validate *validator.Validate) CommentUsecase {
	return &CommentUsecaseImpl{
		Repository: repository,
		DB:         db,
		Validate:   validate,
	}
}
func (usecase *CommentUsecaseImpl) PostComment(ctx context.Context, request web.PostComment) (web.PostComment, error) {
	tx, err := usecase.DB.Begin()
	if err != nil {
		panic(err)
	}

	comment := domain.Comment{
		Message: request.Message,
		PhotoId: request.PhotoId,
		UserId:  request.UserId,
	}

	comment, err = usecase.Repository.PostComment(ctx, tx, comment)
	if err != nil {
		panic(err)
	}

	commentResponse := web.PostComment{
		Id:        comment.Id,
		Message:   comment.Message,
		PhotoId:   comment.PhotoId,
		UserId:    comment.UserId,
		CreatedAt: time.Unix(int64(comment.CreatedAt), 0),
	}

	errr := recover()
	if errr != nil {
		err = tx.Rollback()
		if err != nil {
			panic(err)
		}
	} else {
		err = tx.Commit()
		if err != nil {
			panic(err)
		}
	}

	return commentResponse, nil

}

func (usecase *CommentUsecaseImpl) GetComment(ctx context.Context) ([]web.GetComment, error) {
	tx, err := usecase.DB.Begin()
	if err != nil {
		panic(err)
	}

	comments, user, photo, err := usecase.Repository.GetComment(ctx, tx)
	if err != nil {
		panic(err)
	}
	userWeb := web.UserComment{
		Id:       user.Id,
		Email:    user.Email,
		Username: user.Username,
	}
	photoWeb := web.PhotoComment{
		Id:       photo.Id,
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoUrl: photo.PhotoUrl,
		UserId:   photo.UserId,
	}
	var commentsResponse []web.GetComment
	for _, comment := range comments {
		comment := web.GetComment{
			Id:        comment.Id,
			Message:   comment.Message,
			PhotoId:   comment.PhotoId,
			UserId:    comment.UserId,
			UpdatedAt: time.Unix(int64(comment.UpdatedAt), 0),
			CreatedAt: time.Unix(int64(comment.CreatedAt), 0),
			User:      userWeb,
			Photo:     photoWeb,
		}
		commentsResponse = append(commentsResponse, comment)
	}
	errr := recover()
	if errr != nil {
		err = tx.Rollback()
		if err != nil {
			panic(err)
		}
	} else {
		err = tx.Commit()
		if err != nil {
			panic(err)
		}
	}

	return commentsResponse, nil
}

func (usecase *CommentUsecaseImpl) UpdateComment(ctx context.Context, request web.UpdateComment) (web.UpdateComment, error) {
	tx, err := usecase.DB.Begin()
	if err != nil {
		panic(err)
	}
	comment := domain.Comment{
		Id:      request.Id,
		Message: request.Message,
		UserId:  request.UserId,
	}
	comment, err = usecase.Repository.UpdateComment(ctx, tx, comment)
	if err != nil {
		panic(err)
	}

	commentResponse := web.UpdateComment{
		Id:        comment.Id,
		Message:   comment.Message,
		PhotoId:   comment.PhotoId,
		UserId:    comment.UserId,
		UpdatedAt: time.Unix(int64(comment.UpdatedAt), 0),
	}

	errr := recover()
	if errr != nil {
		err = tx.Rollback()
		if err != nil {
			panic(err)
		}
	} else {
		err = tx.Commit()
		if err != nil {
			panic(err)
		}
	}
	return commentResponse, nil
}

func (usecase *CommentUsecaseImpl) DeleteComment(ctx context.Context, id int) error {
	tx, err := usecase.DB.Begin()
	if err != nil {
		panic(err)
	}

	err = usecase.Repository.DeleteComment(ctx, tx, id)
	if err != nil {
		panic(err)
	}

	errr := recover()
	if errr != nil {
		err = tx.Rollback()
		if err != nil {
			panic(err)
		}
	} else {
		err = tx.Commit()
		if err != nil {
			panic(err)
		}
	}

	return nil
}
