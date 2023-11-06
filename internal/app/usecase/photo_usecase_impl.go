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

type PhotoUsecaseImpl struct {
	Repository repository.PhotoRepository
	Validate   *validator.Validate
	Timeout    int
}

func NewPhotoUsecase(repository repository.PhotoRepository, validate *validator.Validate, timeout int) PhotoUsecase {
	return &PhotoUsecaseImpl{
		Repository: repository,
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

	photo := domain.Photo{
		Title:       request.Title,
		Caption:     request.Caption,
		PhotoBase64: request.PhotoBase64,
		UserId:      request.UserId,
	}

	photo, err = usecase.Repository.PostPhoto(ctx, photo)
	if err != nil {
		return response.PostPhoto{}, err
	}

	tCreate := time.Unix(int64(photo.CreatedAt), 0)

	photoResponse := response.PostPhoto{
		Id:          photo.Id,
		Title:       photo.Title,
		Caption:     photo.Caption,
		PhotoBase64: photo.PhotoBase64,
		UserId:      photo.UserId,
		CreatedAt:   tCreate,
	}

	return photoResponse, nil
}

func (usecase *PhotoUsecaseImpl) GetPhoto(ctx context.Context) ([]response.GetPhoto, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(usecase.Timeout)*time.Second)
	defer cancel()

	photos, users, likes, err := usecase.Repository.GetPhoto(ctx)
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
					Username:       u.Username,
					Email:          u.Email,
					ProfilePicture: u.ProfilePicture,
				}
				break
			}
		}

		var like response.Likes
		for _, l := range likes {
			if l.PhotoId == photo.Id {
				like = response.Likes{
					LikeCount: l.LikeCount,
				}
			}
		}

		photoResp := response.GetPhoto{
			Id:          photo.Id,
			Title:       photo.Title,
			Caption:     photo.Caption,
			PhotoBase64: photo.PhotoBase64,
			UserId:      photo.UserId,
			CreatedAt:   tCreate,
			UpdatedAt:   tUpdate,
			User:        user,
			Likes:       like,
		}

		photoResponse = append(photoResponse, photoResp)
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

	photo := domain.Photo{
		Id:          request.Id,
		Title:       request.Title,
		Caption:     request.Caption,
		PhotoBase64: request.PhotoBase64,
		UserId:      request.UserId,
	}

	photo, err = usecase.Repository.UpdatePhoto(ctx, photo)
	if err != nil {
		return response.UpdatePhoto{}, err
	}

	photoResponse := response.UpdatePhoto{
		Id:          photo.Id,
		Title:       photo.Title,
		Caption:     photo.Caption,
		PhotoBase64: photo.PhotoBase64,
		UserId:      photo.UserId,
		UpdatedAt:   time.Unix(int64(photo.UpdatedAt), 0),
		CreatedAt:   time.Unix(int64(photo.CreatedAt), 0),
	}

	return photoResponse, nil
}

func (usecase *PhotoUsecaseImpl) DeletePhoto(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(usecase.Timeout)*time.Second)
	defer cancel()

	err := usecase.Repository.DeletePhoto(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
