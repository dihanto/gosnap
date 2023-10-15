package usecase

import (
	"context"
	"strconv"
	"time"

	"github.com/dihanto/gosnap/internal/app/repository"
	"github.com/dihanto/gosnap/model/domain"
	"github.com/dihanto/gosnap/model/web/request"
	"github.com/dihanto/gosnap/model/web/response"
	"github.com/go-playground/validator/v10"
)

type LikeUsecaseImpl struct {
	Repository repository.LikeRepository
	Validate   *validator.Validate
	Timeout    int
}

func NewLikeUsecaseImpl(repository repository.LikeRepository, validate *validator.Validate, timeout int) LikeUsecase {
	return &LikeUsecaseImpl{
		Repository: repository,
		Validate:   validate,
		Timeout:    timeout,
	}
}

func (usecase *LikeUsecaseImpl) LikePhoto(ctx context.Context, request request.Like) (response.Like, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(usecase.Timeout)*time.Second)
	defer cancel()

	err := usecase.Validate.Struct(request)
	if err != nil {
		return response.Like{}, err
	}
	photoIdString := strconv.Itoa(request.PhotoId)
	err = usecase.Validate.Var(request.UserId, "likes="+photoIdString)
	if err != nil {
		return response.Like{}, err
	}

	likeRequest := domain.Like{
		PhotoId:   request.PhotoId,
		UserId:    request.UserId,
		CreatedAt: int32(time.Now().Unix()),
	}

	like, err := usecase.Repository.LikePhoto(ctx, likeRequest)
	if err != nil {
		return response.Like{}, err
	}

	likeResponse := response.Like{
		PhotoId:   like.PhotoId,
		UserId:    like.UserId,
		LikeCount: like.LikeCount,
		LikedAt:   time.Now(),
	}

	return likeResponse, nil
}

func (usecase *LikeUsecaseImpl) UnlikePhoto(ctx context.Context, request request.Like) (response.Unlike, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(usecase.Timeout)*time.Second)
	defer cancel()

	err := usecase.Validate.Struct(request)
	if err != nil {
		return response.Unlike{}, err
	}

	likeRequest := domain.Like{
		PhotoId: request.PhotoId,
		UserId:  request.UserId,
	}

	like, err := usecase.Repository.UnlikePhoto(ctx, likeRequest)
	if err != nil {
		return response.Unlike{}, err
	}

	likeResponse := response.Unlike{
		LikeCount: like.LikeCount,
		PhotoId:   like.PhotoId,
	}

	return likeResponse, nil
}

func (usecase *LikeUsecaseImpl) IsLikePhoto(ctx context.Context, request request.Like) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(usecase.Timeout)*time.Second)
	defer cancel()

	err := usecase.Validate.Struct(request)
	if err != nil {
		return false, err
	}

	userIds, err := usecase.Repository.IsLikePhoto(ctx, request.PhotoId)
	if err != nil {
		return false, err
	}

	for _, userId := range userIds {
		if userId == request.UserId {
			return true, nil
		}
	}
	return false, nil
}
