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
		panic(err)
	}
	tx, err := usecase.DB.Begin()
	if err != nil {
		panic(err)
	}
	user := domain.User{
		Email:    request.Email,
		Username: request.Username,
		Password: request.Password,
		Age:      request.Age,
	}
	user, err = usecase.UserRepository.UserRegister(ctx, tx, user)
	if err != nil {
		panic(err)
	}
	tCreate := time.Unix(int64(user.CreatedAt), 0)
	userResponse := web.UserRegister{
		Email:     user.Email,
		Username:  user.Username,
		Password:  user.Password,
		Age:       user.Age,
		CreatedAT: tCreate,
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

	return userResponse, nil

}

func (usecase *UserUsecaseImpl) UserLogin(ctx context.Context, username string, password string) (bool, int, error) {
	tx, err := usecase.DB.Begin()
	if err != nil {
		panic(err)
	}

	response, id, err := usecase.UserRepository.UserLogin(ctx, tx, username, password)
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

	return response, id, nil
}

func (usecase *UserUsecaseImpl) UserUpdate(ctx context.Context, request web.UserUpdate) (web.UserUpdate, error) {
	err := usecase.Validate.Struct(request)
	if err != nil {
		panic(err)
	}
	tx, err := usecase.DB.Begin()
	if err != nil {
		panic(err)
	}

	userReq := domain.User{
		Id:       request.Id,
		Email:    request.Email,
		Username: request.Username,
		Age:      request.Age,
	}

	userResponse, err := usecase.UserRepository.UserUpdate(ctx, tx, userReq)
	if err != nil {
		panic(err)
	}

	tUpdate := time.Unix(int64(userResponse.UpdatedAt), 0)
	log.Println(tUpdate)
	user := web.UserUpdate{
		Id:        userResponse.Id,
		Email:     userResponse.Email,
		Username:  userResponse.Username,
		Age:       userResponse.Age,
		UpdatedAt: tUpdate,
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

	return user, nil

}

func (usecase *UserUsecaseImpl) UserDelete(ctx context.Context, id int) error {
	tx, err := usecase.DB.Begin()
	if err != nil {
		panic(err)
	}

	err = usecase.UserRepository.UserDelete(ctx, tx, id)
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
