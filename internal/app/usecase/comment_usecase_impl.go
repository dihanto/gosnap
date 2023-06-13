package usecase

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/dihanto/gosnap/internal/app/repository"
	"github.com/dihanto/gosnap/model/domain"
	"github.com/dihanto/gosnap/model/web/request"
	"github.com/dihanto/gosnap/model/web/response"
	"github.com/go-playground/validator/v10"
)

type CommentUsecaseImpl struct {
	Repository repository.CommentRepository
	DB         *sql.DB
	Validate   *validator.Validate
	Timeout    int
}

func NewCommentUsecase(repository repository.CommentRepository, db *sql.DB, validate *validator.Validate, timeout int) CommentUsecase {
	return &CommentUsecaseImpl{
		Repository: repository,
		DB:         db,
		Validate:   validate,
		Timeout:    timeout,
	}
}
func (usecase *CommentUsecaseImpl) PostComment(c context.Context, request request.Comment) (response.PostComment, error) {
	ctx, cancel := context.WithTimeout(c, time.Duration(usecase.Timeout)*time.Second)
	defer cancel()

	err := usecase.Validate.Struct(request)
	if err != nil {
		return response.PostComment{}, err
	}

	tx, err := usecase.DB.Begin()
	if err != nil {
		return response.PostComment{}, err
	}
	defer func() {
		if r := recover(); r != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				log.Println("Failed to rollback transaction:", rollbackErr)
			}
			panic(r)
		}
	}()

	comment := domain.Comment{
		Message: request.Message,
		PhotoId: request.PhotoId,
		UserId:  request.UserId,
	}

	comment, err = usecase.Repository.PostComment(ctx, tx, comment)
	if err != nil {
		return response.PostComment{}, err
	}

	err = tx.Commit()
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			log.Println("Failed to rollback transaction:", rollbackErr)
		}
		return response.PostComment{}, err
	}

	commentResponse := response.PostComment{
		Id:        comment.Id,
		Message:   comment.Message,
		PhotoId:   comment.PhotoId,
		UserId:    comment.UserId,
		CreatedAt: time.Unix(int64(comment.CreatedAt), 0),
	}

	return commentResponse, nil
}

func (usecase *CommentUsecaseImpl) GetComment(c context.Context) ([]response.GetComment, error) {
	ctx, cancel := context.WithTimeout(c, time.Duration(usecase.Timeout)*time.Second)
	defer cancel()

	tx, err := usecase.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				log.Println("Failed to rollback transaction:", rollbackErr)
			}
			panic(r)
		}
	}()

	comments, users, photos, err := usecase.Repository.GetComment(ctx, tx)
	if err != nil {
		return nil, err
	}

	var commentsResponse []response.GetComment
	for _, comment := range comments {

		var userResponse response.UserComment
		for _, user := range users {
			if user.Id == comment.UserId {
				userResponse = response.UserComment{
					Id:       user.Id,
					Username: user.Username,
					Email:    user.Email,
				}
			}
		}
		var photoResponse response.PhotoComment
		for _, photo := range photos {
			if photo.Id == comment.PhotoId {
				photoResponse = response.PhotoComment{
					Id:       photo.Id,
					Title:    photo.Title,
					Caption:  photo.Caption,
					PhotoUrl: photo.PhotoUrl,
					UserId:   photo.UserId,
				}
			}
		}

		commentResponse := response.GetComment{
			Id:        comment.Id,
			Message:   comment.Message,
			PhotoId:   comment.PhotoId,
			UserId:    comment.UserId,
			UpdatedAt: time.Unix(int64(comment.UpdatedAt), 0),
			CreatedAt: time.Unix(int64(comment.CreatedAt), 0),
			User:      userResponse,
			Photo:     photoResponse,
		}
		commentsResponse = append(commentsResponse, commentResponse)
	}

	err = tx.Commit()
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			log.Println("Failed to rollback transaction:", rollbackErr)
		}
		return nil, err
	}

	return commentsResponse, nil
}

func (usecase *CommentUsecaseImpl) UpdateComment(c context.Context, request request.Comment) (response.UpdateComment, error) {
	ctx, cancel := context.WithTimeout(c, time.Duration(usecase.Timeout)*time.Second)
	defer cancel()

	err := usecase.Validate.Struct(request)
	if err != nil {
		return response.UpdateComment{}, err
	}

	tx, err := usecase.DB.Begin()
	if err != nil {
		return response.UpdateComment{}, err
	}
	defer func() {
		if r := recover(); r != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				log.Println("Failed to rollback transaction:", rollbackErr)
			}
			panic(r)
		}
	}()

	comment := domain.Comment{
		Id:      request.Id,
		Message: request.Message,
		UserId:  request.UserId,
	}

	comment, err = usecase.Repository.UpdateComment(ctx, tx, comment)
	if err != nil {
		return response.UpdateComment{}, err
	}

	commentResponse := response.UpdateComment{
		Id:        comment.Id,
		Message:   comment.Message,
		PhotoId:   comment.PhotoId,
		UserId:    comment.UserId,
		UpdatedAt: time.Unix(int64(comment.UpdatedAt), 0),
	}

	err = tx.Commit()
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			log.Println("Failed to rollback transaction:", rollbackErr)
		}
		return response.UpdateComment{}, err
	}

	return commentResponse, nil
}

func (usecase *CommentUsecaseImpl) DeleteComment(c context.Context, id int) error {
	ctx, cancel := context.WithTimeout(c, time.Duration(usecase.Timeout)*time.Second)
	defer cancel()

	tx, err := usecase.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				log.Println("Failed to rollback transaction:", rollbackErr)
			}
			panic(r)
		}
	}()

	err = usecase.Repository.DeleteComment(ctx, tx, id)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			log.Println("Failed to rollback transaction:", rollbackErr)
		}
		return err
	}

	return nil
}
