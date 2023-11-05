package usecase

import (
	"context"
	"time"

	"github.com/dihanto/gosnap/internal/app/repository"
	"github.com/dihanto/gosnap/model/domain"
	"github.com/dihanto/gosnap/model/web/request"
	"github.com/dihanto/gosnap/model/web/response"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type UserUsecaseImpl struct {
	Repository repository.UserRepository
	Validate   *validator.Validate
	Timeout    int
}

func NewUserUsecase(repository repository.UserRepository, validate *validator.Validate, timeout int) UserUsecase {
	return &UserUsecaseImpl{
		Repository: repository,
		Validate:   validate,
		Timeout:    timeout,
	}
}
func (usecase *UserUsecaseImpl) UserRegister(ctx context.Context, request request.UserRegister) (response.UserRegister, error) {
	err := usecase.Validate.Struct(request)
	if err != nil {
		return response.UserRegister{}, err
	}

	ctx, cancel := context.WithTimeout(ctx, time.Duration(usecase.Timeout)*time.Second)
	defer cancel()

	user := domain.User{
		Email:    request.Email,
		Username: request.Username,
		Name:     request.Name,
		Password: request.Password,
		Age:      request.Age,
	}
	user.Id = uuid.New()

	user, err = usecase.Repository.UserRegister(ctx, user)
	if err != nil {
		return response.UserRegister{}, err
	}

	tCreate := time.Unix(int64(user.CreatedAt), 0)
	userResponse := response.UserRegister{
		Email:     user.Email,
		Username:  user.Username,
		Password:  user.Password,
		Age:       user.Age,
		CreatedAT: tCreate,
	}

	return userResponse, nil
}

func (usecase *UserUsecaseImpl) UserLogin(ctx context.Context, username string, password string) (bool, uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(usecase.Timeout)*time.Second)
	defer cancel()

	response, id, err := usecase.Repository.UserLogin(ctx, username, password)
	if err != nil {
		return false, uuid.Nil, err
	}

	return response, id, nil
}

func (usecase *UserUsecaseImpl) UserUpdate(ctx context.Context, request request.UserUpdate) (response.UserUpdate, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(usecase.Timeout)*time.Second)
	defer cancel()

	err := usecase.Validate.Struct(request)
	if err != nil {
		return response.UserUpdate{}, err
	}

	userReq := domain.User{
		Id:             request.Id,
		Username:       request.Username,
		Email:          request.Email,
		ProfilePicture: request.ProfilePicture,
	}

	userResponse, err := usecase.Repository.UserUpdate(ctx, userReq)
	if err != nil {
		return response.UserUpdate{}, err
	}

	tUpdate := time.Unix(int64(userResponse.UpdatedAt), 0)
	user := response.UserUpdate{
		Id:             userResponse.Id,
		Email:          userResponse.Email,
		Username:       userResponse.Username,
		Age:            userResponse.Age,
		ProfilePicture: userResponse.ProfilePicture,
		UpdatedAt:      tUpdate,
	}

	return user, nil
}

func (usecase *UserUsecaseImpl) UserDelete(ctx context.Context, id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(usecase.Timeout)*time.Second)
	defer cancel()

	err := usecase.Repository.UserDelete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *UserUsecaseImpl) FindUser(ctx context.Context, id uuid.UUID) (response.FindUser, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(usecase.Timeout)*time.Second)
	defer cancel()

	err := usecase.Validate.Var(id, "required")
	if err != nil {
		return response.FindUser{}, err
	}

	user, err := usecase.Repository.FindUser(ctx, id)
	if err != nil {
		return response.FindUser{}, nil
	}

	userResponse := response.FindUser{
		Username:       user.Username,
		Name:           user.Name,
		ProfilePicture: user.ProfilePicture,
	}

	return userResponse, nil
}

func (usecase *UserUsecaseImpl) FindAllUser(ctx context.Context, username string) (users []response.FindAllUser, err error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(usecase.Timeout)*time.Second)
	defer cancel()

	usersRepository, err := usecase.Repository.FindAllUser(ctx)
	if err != nil {
		return
	}

	for _, userRepository := range usersRepository {
		if username != userRepository.Username {
			user := response.FindAllUser{
				Username: userRepository.Username,
			}
			users = append(users, user)
		}
	}

	return
}
