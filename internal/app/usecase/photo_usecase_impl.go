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

type PhotoUsecaseImpl struct {
	Repository repository.PhotoRepository
	DB         *sql.DB
	Validate   *validator.Validate
	Timeout    int
}

func NewPhotoUsecase(repository repository.PhotoRepository, db *sql.DB, validate *validator.Validate, timeout int) PhotoUsecase {
	return &PhotoUsecaseImpl{
		Repository: repository,
		DB:         db,
		Validate:   validate,
		Timeout:    timeout,
	}
}

func (usecase *PhotoUsecaseImpl) PostPhoto(ctx context.Context, request request.Photo) (response.PostPhoto, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(usecase.Timeout)*time.Second)
	defer cancel()

	err := usecase.Validate.Struct(request)
	if err != nil {
		return response.PostPhoto{}, err
	}

	tx, err := usecase.DB.Begin()
	if err != nil {
		return response.PostPhoto{}, err
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

	photo := domain.Photo{
		Title:    request.Title,
		Caption:  request.Caption,
		PhotoUrl: request.PhotoUrl,
		UserId:   request.UserId,
	}

	photo, err = usecase.Repository.PostPhoto(ctx, tx, photo)
	if err != nil {
		return response.PostPhoto{}, err
	}

	err = tx.Commit()
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			log.Println("Failed to rollback transaction:", rollbackErr)
		}
		return response.PostPhoto{}, err
	}

	tCreate := time.Unix(int64(photo.CreatedAt), 0)

	photoResponse := response.PostPhoto{
		Id:        photo.Id,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoUrl:  photo.PhotoUrl,
		UserId:    photo.UserId,
		CreatedAt: tCreate,
	}

	return photoResponse, nil
}

func (usecase *PhotoUsecaseImpl) GetPhoto(ctx context.Context) ([]response.GetPhoto, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(usecase.Timeout)*time.Second)
	defer cancel()

	tx, err := usecase.DB.Begin()
	if err != nil {
		return nil, err
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

	photos, users, err := usecase.Repository.GetPhoto(ctx, tx)
	if err != nil {
		return nil, err
	}

	var photoResponse []response.GetPhoto
	for _, photo := range photos {
		tCreate := time.Unix(int64(photo.CreatedAt), 0)
		tUpdate := time.Unix(int64(photo.UpdatedAt), 0)

		var user response.User
		for _, u := range users {
			if u.Id == photo.UserId {
				user = response.User{
					Username: u.Username,
					Email:    u.Email,
				}
				break
			}
		}

		photoResp := response.GetPhoto{
			Id:        photo.Id,
			Title:     photo.Title,
			Caption:   photo.Caption,
			PhotoUrl:  photo.PhotoUrl,
			UserId:    photo.UserId,
			CreatedAt: tCreate,
			UpdatedAt: tUpdate,
			User:      user,
		}

		photoResponse = append(photoResponse, photoResp)
	}

	err = tx.Commit()
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			log.Println("Failed to rollback transaction:", rollbackErr)
		}
		return nil, err
	}

	return photoResponse, nil
}

func (usecase *PhotoUsecaseImpl) UpdatePhoto(ctx context.Context, request request.Photo) (response.UpdatePhoto, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(usecase.Timeout)*time.Second)
	defer cancel()

	err := usecase.Validate.Struct(request)
	if err != nil {
		return response.UpdatePhoto{}, err
	}

	tx, err := usecase.DB.Begin()
	if err != nil {
		return response.UpdatePhoto{}, err
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

	photo := domain.Photo{
		Id:       request.Id,
		Title:    request.Title,
		Caption:  request.Caption,
		PhotoUrl: request.PhotoUrl,
		UserId:   request.UserId,
	}

	photo, err = usecase.Repository.UpdatePhoto(ctx, tx, photo)
	if err != nil {
		return response.UpdatePhoto{}, err
	}

	photoResponse := response.UpdatePhoto{
		Id:        photo.Id,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoUrl:  photo.PhotoUrl,
		UserId:    photo.UserId,
		UpdatedAt: time.Unix(int64(photo.UpdatedAt), 0),
		CreatedAt: time.Unix(int64(photo.CreatedAt), 0),
	}

	err = tx.Commit()
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			log.Println("Failed to rollback transaction:", rollbackErr)
		}
		return response.UpdatePhoto{}, err
	}

	return photoResponse, nil
}

func (usecase *PhotoUsecaseImpl) DeletePhoto(ctx context.Context, id int) error {
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

	err = usecase.Repository.DeletePhoto(ctx, tx, id)
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
