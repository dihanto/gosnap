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

type UserUsecaseImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserUsecase(userRepository repository.UserRepository, db *sql.DB, validate *validator.Validate) UserUsecase {
	return &UserUsecaseImpl{
		UserRepository: userRepository,
		DB:             db,
		Validate:       validate,
	}
}
func (usecase *UserUsecaseImpl) UserRegister(ctx context.Context, request web.UserRegister) (web.UserRegister, error) {
	err := usecase.Validate.Struct(request)
	if err != nil {
		return web.UserRegister{}, err
	}

	tx, err := usecase.DB.Begin()
	if err != nil {
		return web.UserRegister{}, err
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

	user := domain.User{
		Email:    request.Email,
		Username: request.Username,
		Password: request.Password,
		Age:      request.Age,
	}
	user, err = usecase.UserRepository.UserRegister(ctx, tx, user)
	if err != nil {
		return web.UserRegister{}, err
	}

	err = tx.Commit()
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			log.Println("Failed to rollback transaction:", rollbackErr)
		}
		return web.UserRegister{}, err
	}

	tCreate := time.Unix(int64(user.CreatedAt), 0)
	userResponse := web.UserRegister{
		Email:     user.Email,
		Username:  user.Username,
		Password:  user.Password,
		Age:       user.Age,
		CreatedAT: tCreate,
	}

	return userResponse, nil
}

func (usecase *UserUsecaseImpl) UserLogin(ctx context.Context, username string, password string) (bool, int, error) {
	tx, err := usecase.DB.Begin()
	if err != nil {
		return false, 0, err
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

	response, id, err := usecase.UserRepository.UserLogin(ctx, tx, username, password)
	if err != nil {
		return false, 0, err
	}

	err = tx.Commit()
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			log.Println("Failed to rollback transaction:", rollbackErr)
		}
		return false, 0, err
	}

	return response, id, nil
}

func (usecase *UserUsecaseImpl) UserUpdate(ctx context.Context, request web.UserUpdate) (web.UserUpdate, error) {
	err := usecase.Validate.Struct(request)
	if err != nil {
		return web.UserUpdate{}, err
	}

	tx, err := usecase.DB.Begin()
	if err != nil {
		return web.UserUpdate{}, err
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

	userReq := domain.User{
		Id:       request.Id,
		Email:    request.Email,
		Username: request.Username,
		Age:      request.Age,
	}

	userResponse, err := usecase.UserRepository.UserUpdate(ctx, tx, userReq)
	if err != nil {
		return web.UserUpdate{}, err
	}

	err = tx.Commit()
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			log.Println("Failed to rollback transaction:", rollbackErr)
		}
		return web.UserUpdate{}, err
	}

	tUpdate := time.Unix(int64(userResponse.UpdatedAt), 0)
	user := web.UserUpdate{
		Id:        userResponse.Id,
		Email:     userResponse.Email,
		Username:  userResponse.Username,
		Age:       userResponse.Age,
		UpdatedAt: tUpdate,
	}

	return user, nil
}

func (usecase *UserUsecaseImpl) UserDelete(ctx context.Context, id int) error {
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
