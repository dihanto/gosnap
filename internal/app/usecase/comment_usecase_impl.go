package usecase

import (
	"context"
	"database/sql"
	"log"
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
	Timeout    int
}

func NewCommentUsecase(repository repository.CommentRepository, db *sql.DB, validate *validator.Validate) CommentUsecase {
	return &CommentUsecaseImpl{
		Repository: repository,
		DB:         db,
		Validate:   validate,
	}
}
func (usecase *CommentUsecaseImpl) PostComment(ctx context.Context, request web.PostComment) (web.PostComment, error) {
	// Validate the request
	err := usecase.Validate.Struct(request)
	if err != nil {
		return web.PostComment{}, err
	}

	// Begin transaction
	tx, err := usecase.DB.Begin()
	if err != nil {
		return web.PostComment{}, err
	}
	defer func() {
		if r := recover(); r != nil {
			// Rollback transaction on panic
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				// Handle rollback error if necessary
				log.Println("Failed to rollback transaction:", rollbackErr)
			}
			panic(r) // Re-panic the recovered panic
		}
	}()

	// Create the comment
	comment := domain.Comment{
		Message: request.Message,
		PhotoId: request.PhotoId,
		UserId:  request.UserId,
	}
	comment, err = usecase.Repository.PostComment(ctx, tx, comment)
	if err != nil {
		return web.PostComment{}, err
	}

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		// Rollback on commit error
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			// Handle rollback error if necessary
			log.Println("Failed to rollback transaction:", rollbackErr)
		}
		return web.PostComment{}, err
	}

	// Prepare response
	commentResponse := web.PostComment{
		Id:        comment.Id,
		Message:   comment.Message,
		PhotoId:   comment.PhotoId,
		UserId:    comment.UserId,
		CreatedAt: time.Unix(int64(comment.CreatedAt), 0),
	}

	return commentResponse, nil
}

func (usecase *CommentUsecaseImpl) GetComment(ctx context.Context) ([]web.GetComment, error) {
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

	comments, user, photo, err := usecase.Repository.GetComment(ctx, tx)
	if err != nil {
		return nil, err
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
		commentResponse := web.GetComment{
			Id:        comment.Id,
			Message:   comment.Message,
			PhotoId:   comment.PhotoId,
			UserId:    comment.UserId,
			UpdatedAt: time.Unix(int64(comment.UpdatedAt), 0),
			CreatedAt: time.Unix(int64(comment.CreatedAt), 0),
			User:      userWeb,
			Photo:     photoWeb,
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

func (usecase *CommentUsecaseImpl) UpdateComment(ctx context.Context, request web.UpdateComment) (web.UpdateComment, error) {
	err := usecase.Validate.Struct(request)
	if err != nil {
		return web.UpdateComment{}, err
	}

	tx, err := usecase.DB.Begin()
	if err != nil {
		return web.UpdateComment{}, err
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
		return web.UpdateComment{}, err
	}

	commentResponse := web.UpdateComment{
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
		return web.UpdateComment{}, err
	}

	return commentResponse, nil
}

func (usecase *CommentUsecaseImpl) DeleteComment(ctx context.Context, id int) error {
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
