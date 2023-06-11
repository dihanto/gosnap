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

type PhotoUsecaseImpl struct {
	Repository repository.PhotoRepository
	DB         *sql.DB
	Validate   *validator.Validate
}

func NewPhotoUsecase(repository repository.PhotoRepository, db *sql.DB, validate *validator.Validate) PhotoUsecase {
	return &PhotoUsecaseImpl{
		Repository: repository,
		DB:         db,
		Validate:   validate,
	}
}

func (usecase *PhotoUsecaseImpl) PostPhoto(ctx context.Context, request web.Photo) (web.Photo, error) {
	err := usecase.Validate.Struct(request)
	if err != nil {
		return web.Photo{}, err
	}

	tx, err := usecase.DB.Begin()
	if err != nil {
		return web.Photo{}, err
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

	photo := domain.Photo{
		Title:    request.Title,
		Caption:  request.Caption,
		PhotoUrl: request.PhotoUrl,
		UserId:   request.UserId,
	}

	photo, err = usecase.Repository.PostPhoto(ctx, tx, photo)
	if err != nil {
		return web.Photo{}, err
	}

	err = tx.Commit()
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			log.Println("Failed to rollback transaction:", rollbackErr)
		}
		return web.Photo{}, err
	}

	tCreate := time.Unix(int64(photo.CreatedAt), 0)

	photoResponse := web.Photo{
		Id:        photo.Id,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoUrl:  photo.PhotoUrl,
		UserId:    photo.UserId,
		CreatedAt: tCreate,
	}

	return photoResponse, nil
}

func (usecase *PhotoUsecaseImpl) GetPhoto(ctx context.Context) ([]web.GetPhoto, error) {
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

	photos, user, err := usecase.Repository.GetPhoto(ctx, tx)
	if err != nil {
		return nil, err
	}

	userWeb := web.User{
		Username: user.Username,
		Email:    user.Email,
	}

	var photoResponse []web.GetPhoto
	for _, photo := range photos {
		tCreate := time.Unix(int64(photo.CreatedAt), 0)
		tUpdate := time.Unix(int64(photo.UpdatedAt), 0)
		photo := web.GetPhoto{
			Id:        photo.Id,
			Title:     photo.Title,
			Caption:   photo.Caption,
			PhotoUrl:  photo.PhotoUrl,
			UserId:    photo.UserId,
			CreatedAt: tCreate,
			UpdatedAt: tUpdate,
			User:      userWeb,
		}
		photoResponse = append(photoResponse, photo)
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

func (usecase *PhotoUsecaseImpl) UpdatePhoto(ctx context.Context, request web.Photo) (web.UpdatePhoto, error) {
	err := usecase.Validate.Struct(request)
	if err != nil {
		return web.UpdatePhoto{}, err
	}

	tx, err := usecase.DB.Begin()
	if err != nil {
		return web.UpdatePhoto{}, err
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

	photo := domain.Photo{
		Id:       request.Id,
		Title:    request.Title,
		Caption:  request.Caption,
		PhotoUrl: request.PhotoUrl,
		UserId:   request.UserId,
	}

	photo, err = usecase.Repository.UpdatePhoto(ctx, tx, photo)
	if err != nil {
		return web.UpdatePhoto{}, err
	}

	photoResponse := web.UpdatePhoto{
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
		return web.UpdatePhoto{}, err
	}

	return photoResponse, nil
}

func (usecase *PhotoUsecaseImpl) DeletePhoto(ctx context.Context, id int) error {
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
