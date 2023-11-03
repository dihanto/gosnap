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

type FollowUsecaseImpl struct {
	Repository repository.FollowRepository
	Validate   *validator.Validate
	Timeout    int
}

func NewFollowUsecaseImpl(repository repository.FollowRepository, validate *validator.Validate, timeout int) FollowUsecase {
	return &FollowUsecaseImpl{
		Repository: repository,
		Validate:   validate,
		Timeout:    timeout,
	}
}

func (usecase *FollowUsecaseImpl) FollowUser(ctx context.Context, request request.Follow) (response.Follow, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(usecase.Timeout)*time.Second)
	defer cancel()

	err := usecase.Validate.Struct(request)
	if err != nil {
		return response.Follow{}, err
	}
	err = usecase.Validate.Var(request.FollowerUsername, "follow="+request.TargetUsername)
	if err != nil {
		return response.Follow{}, err
	}

	followRequest := domain.Follow{
		FollowerUsername: request.FollowerUsername,
		TargetUsername:   request.TargetUsername,
	}

	follow, err := usecase.Repository.FollowUser(ctx, followRequest)
	if err != nil {
		return response.Follow{}, err
	}

	followResponse := response.Follow{
		FollowerCount: follow.FollowerCount,
	}

	return followResponse, nil
}

func (usecase *FollowUsecaseImpl) UnFollowUser(ctx context.Context, request request.Follow) (response.Follow, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(usecase.Timeout)*time.Second)
	defer cancel()

	err := usecase.Validate.Struct(request)
	if err != nil {
		return response.Follow{}, err
	}

	followRequest := domain.Follow{
		FollowerUsername: request.FollowerUsername,
		TargetUsername:   request.TargetUsername,
	}

	follow, err := usecase.Repository.UnFollowUser(ctx, followRequest)
	if err != nil {
		return response.Follow{}, err
	}

	followResponse := response.Follow{
		FollowerCount: follow.FollowerCount,
	}

	return followResponse, nil

}

func (usecase *FollowUsecaseImpl) GetFollower(ctx context.Context, request request.Follow) (followers response.GetFollower, err error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(usecase.Timeout)*time.Second)
	defer cancel()

	err = usecase.Validate.Var(request.TargetUsername, "required")
	if err != nil {
		return
	}
	followersRequest := domain.Follow{
		TargetUsername: request.TargetUsername,
	}

	followersResponse, err := usecase.Repository.GetFollower(ctx, followersRequest)
	if err != nil {
		return
	}

	for _, followerResponse := range followersResponse {
		followers.Username = append(followers.Username, followerResponse.Username)
	}
	followers.FollowerCount = len(followers.Username)

	return
}

func (usecase *FollowUsecaseImpl) GetFollowing(ctx context.Context, request request.Follow) (follows response.GetFollowing, err error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(usecase.Timeout)*time.Second)
	defer cancel()

	err = usecase.Validate.Var(request.TargetUsername, "required")
	if err != nil {
		return
	}

	followRequest := domain.Follow{
		TargetUsername: request.TargetUsername,
	}

	followResponses, err := usecase.Repository.GetFollowing(ctx, followRequest)
	if err != nil {
		return
	}

	for _, followResponse := range followResponses {
		follows.Username = append(follows.Username, followResponse.Username)
	}
	follows.FollowingCount = len(follows.Username)

	return
}
