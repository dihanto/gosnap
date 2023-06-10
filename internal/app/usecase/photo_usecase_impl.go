package usecase

import (
	"context"
	"database/sql"
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
		panic(err)
	}
	tx, err := usecase.DB.Begin()
	if err != nil {
		panic(err)
	}

	photo := domain.Photo{
		Title:    request.Title,
		Caption:  request.Caption,
		PhotoUrl: request.PhotoUrl,
		UserId:   request.UserId,
	}

	photo, err = usecase.Repository.PostPhoto(ctx, tx, photo)
	if err != nil {
		panic(err)
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

	return photoResponse, nil
}

func (usecase *PhotoUsecaseImpl) GetPhoto(ctx context.Context) ([]web.GetPhoto, error) {
	tx, err := usecase.DB.Begin()
	if err != nil {
		panic(err)
	}

	photos, user, err := usecase.Repository.GetPhoto(ctx, tx)
	if err != nil {
		panic(err)
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

	return photoResponse, nil
}

func (usecase *PhotoUsecaseImpl) UpdatePhoto(ctx context.Context, request web.Photo) (web.UpdatePhoto, error) {
	err := usecase.Validate.Struct(request)
	if err != nil {
		panic(err)
	}
	tx, err := usecase.DB.Begin()
	if err != nil {
		panic(err)
	}

	photo := domain.Photo{
		Title:    request.Title,
		Caption:  request.Title,
		PhotoUrl: request.PhotoUrl,
		UserId:   request.UserId,
		Id:       request.Id,
	}

	photo, err = usecase.Repository.UpdatePhoto(ctx, tx, photo)
	if err != nil {
		panic(err)
	}

	timeResponse := time.Unix(int64(photo.UpdatedAt), 0)

	photoResponse := web.UpdatePhoto{
		Id:        photo.Id,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoUrl:  photo.PhotoUrl,
		UserId:    photo.UserId,
		UpdatedAt: timeResponse,
		CreatedAt: time.Unix(int64(photo.CreatedAt), 0),
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

	return photoResponse, nil

}

func (usecase *PhotoUsecaseImpl) DeletePhoto(ctx context.Context, id int) error {
	tx, err := usecase.DB.Begin()
	if err != nil {
		panic(err)
	}

	err = usecase.Repository.DeletePhoto(ctx, tx, id)
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
