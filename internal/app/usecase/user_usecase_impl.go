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
	"github.com/google/uuid"
)

type UserUsecaseImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
	Timeout        int
}

func NewUserUsecase(userRepository repository.UserRepository, db *sql.DB, validate *validator.Validate, timeout int) UserUsecase {
	return &UserUsecaseImpl{
		UserRepository: userRepository,
		DB:             db,
		Validate:       validate,
		Timeout:        timeout,
	}
}
func (usecase *UserUsecaseImpl) UserRegister(ctx context.Context, request request.UserRegister) (response.UserRegister, error) {
	err := usecase.Validate.Struct(request)
	if err != nil {
		return response.UserRegister{}, err
	}

	ctx, cancel := context.WithTimeout(ctx, time.Duration(usecase.Timeout)*time.Second)
	defer cancel()

	tx, err := usecase.DB.Begin()
	if err != nil {
		return response.UserRegister{}, err
	}
	defer func() {
		if recover := recover(); recover != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				log.Println("Failed to rollback transaction:", rollbackErr)
			}
			panic(recover)
		}
	}()

	user := domain.User{
		Email:    request.Email,
		Username: request.Username,
		Password: request.Password,
		Age:      request.Age,
	}
	user.Id = uuid.New()

	user, err = usecase.UserRepository.UserRegister(ctx, tx, user)
	if err != nil {
		return response.UserRegister{}, err
	}

	err = tx.Commit()
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			log.Println("Failed to rollback transaction:", rollbackErr)
		}
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

	tx, err := usecase.DB.Begin()
	if err != nil {
		return false, uuid.Nil, err
	}
	defer func() {
		if recover := recover(); recover != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				log.Println("Failed to rollback transaction:", rollbackErr)
			}
			panic(recover)
		}
	}()

	response, id, err := usecase.UserRepository.UserLogin(ctx, tx, username, password)
	if err != nil {
		return false, uuid.Nil, err
	}

	err = tx.Commit()
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			log.Println("Failed to rollback transaction:", rollbackErr)
		}
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

	tx, err := usecase.DB.Begin()
	if err != nil {
		return response.UserUpdate{}, err
	}
	defer func() {
		if recover := recover(); recover != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				log.Println("Failed to rollback transaction:", rollbackErr)
			}
			panic(recover)
		}
	}()

	userReq := domain.User{
		Id:       request.Id,
		Username: request.Username,
		Email:    request.Email,
	}

	userResponse, err := usecase.UserRepository.UserUpdate(ctx, tx, userReq)
	if err != nil {
		return response.UserUpdate{}, err
	}

	err = tx.Commit()
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			log.Println("Failed to rollback transaction:", rollbackErr)
		}
		return response.UserUpdate{}, err
	}

	tUpdate := time.Unix(int64(userResponse.UpdatedAt), 0)
	user := response.UserUpdate{
		Id:        userResponse.Id,
		Email:     userResponse.Email,
		Username:  userResponse.Username,
		Age:       userResponse.Age,
		UpdatedAt: tUpdate,
	}

	return user, nil
}

func (usecase *UserUsecaseImpl) UserDelete(ctx context.Context, id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(usecase.Timeout)*time.Second)
	defer cancel()

	tx, err := usecase.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if recover := recover(); recover != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				log.Println("Failed to rollback transaction:", rollbackErr)
			}
			panic(recover)
		}
	}()

	err = usecase.UserRepository.UserDelete(ctx, tx, id)
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
