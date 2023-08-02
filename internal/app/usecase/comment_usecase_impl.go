package usecase

import (
	"context"
	"time"

	"github.com/dihanto/gosnap/internal/app/repository"
	"github.com/dihanto/gosnap/model/domain"
	"github.com/dihanto/gosnap/model/web/request"
	"github.com/dihanto/gosnap/model/web/response"
	"github.com/go-playground/validator/v10"
)

type CommentUsecaseImpl struct {
	Repository repository.CommentRepository
	Validate   *validator.Validate
	Timeout    int
}

func NewCommentUsecase(repository repository.CommentRepository, validate *validator.Validate, timeout int) CommentUsecase {
	return &CommentUsecaseImpl{
		Repository: repository,
		Validate:   validate,
		Timeout:    timeout,
	}
}
func (usecase *CommentUsecaseImpl) PostComment(ctx context.Context, request request.Comment) (response.PostComment, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(usecase.Timeout)*time.Second)
	defer cancel()

	err := usecase.Validate.Struct(request)
	if err != nil {
		return response.PostComment{}, err
	}

	comment := domain.Comment{
		Message: request.Message,
		PhotoId: request.PhotoId,
		UserId:  request.UserId,
	}

	comment, err = usecase.Repository.PostComment(ctx, comment)
	if err != nil {
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

func (usecase *CommentUsecaseImpl) GetComment(ctx context.Context) ([]response.GetComment, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(usecase.Timeout)*time.Second)
	defer cancel()

	comments, users, photos, err := usecase.Repository.GetComment(ctx)
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

	return commentsResponse, nil
}

func (usecase *CommentUsecaseImpl) UpdateComment(ctx context.Context, request request.Comment) (response.UpdateComment, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(usecase.Timeout)*time.Second)
	defer cancel()

	err := usecase.Validate.Struct(request)
	if err != nil {
		return response.UpdateComment{}, err
	}

	comment := domain.Comment{
		Id:      request.Id,
		Message: request.Message,
		UserId:  request.UserId,
	}

	comment, err = usecase.Repository.UpdateComment(ctx, comment)
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

	return commentResponse, nil
}

func (usecase *CommentUsecaseImpl) DeleteComment(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(usecase.Timeout)*time.Second)
	defer cancel()

	err := usecase.Repository.DeleteComment(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
