package usecase

import (
	"context"
	"time"

	"github.com/dihanto/gosnap/internal/app/repository"
	"github.com/dihanto/gosnap/model/domain"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
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

func (usecase *FollowUsecaseImpl) FollowUser(ctx context.Context, followerId uuid.UUID, username string) (user domain.User, err error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(usecase.Timeout)*time.Second)
	defer cancel()

	err = usecase.Validate.Var(followerId, "required")
	if err != nil {
		return
	}
	err = usecase.Validate.Var(username, "required")
	if err != nil {
		return
	}

	user, err = usecase.Repository.FollowUser(ctx, followerId, username)
	if err != nil {
		return
	}

	return
}

func (usecase *FollowUsecaseImpl) UnFollowUser(ctx context.Context, id int) (err error) {
	panic("not implemented") // TODO: Implement
}
