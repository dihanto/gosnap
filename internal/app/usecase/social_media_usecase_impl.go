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

type SocialMediaUsecaseImpl struct {
	Repository repository.SocialMediaRepository
	DB         *sql.DB
	Validate   *validator.Validate
	Timeout    int
}

func NewSocialMediaUsecase(repository repository.SocialMediaRepository, db *sql.DB, validate *validator.Validate, timeout int) SocialMediaUsecase {
	return &SocialMediaUsecaseImpl{
		Repository: repository,
		DB:         db,
		Validate:   validate,
		Timeout:    timeout,
	}
}

func (usecase *SocialMediaUsecaseImpl) PostSocialMedia(c context.Context, request request.SocialMedia) (response.PostSocialMedia, error) {
	ctx, cancel := context.WithTimeout(c, time.Duration(usecase.Timeout)*time.Second)
	defer cancel()

	err := usecase.Validate.Struct(request)
	if err != nil {
		return response.PostSocialMedia{}, err
	}

	tx, err := usecase.DB.Begin()
	if err != nil {
		return response.PostSocialMedia{}, err
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

	socialMedia := domain.SocialMedia{
		Name:           request.Name,
		SocialMediaUrl: request.SocialMediaUrl,
		UserId:         request.UserId,
	}
	socialMedia.Id = uuid.New()

	socialMedia, err = usecase.Repository.PostSocialMedia(ctx, tx, socialMedia)
	if err != nil {
		return response.PostSocialMedia{}, err
	}

	err = tx.Commit()
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			log.Println("Failed to rollback transaction:", rollbackErr)
		}
		return response.PostSocialMedia{}, err
	}

	socialMediaResponse := response.PostSocialMedia{
		Id:             socialMedia.Id,
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
		UserId:         socialMedia.UserId,
		CreatedAt:      time.Unix(int64(socialMedia.CreatedAt), 0),
	}

	return socialMediaResponse, nil
}

func (usecase *SocialMediaUsecaseImpl) GetSocialMedia(c context.Context) ([]response.GetSocialMedia, error) {
	ctx, cancel := context.WithTimeout(c, time.Duration(usecase.Timeout)*time.Second)
	defer cancel()

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

	socialMedias, users, err := usecase.Repository.GetSocialMedia(ctx, tx)
	if err != nil {
		return nil, err
	}

	var socialMediasResponse []response.GetSocialMedia
	for _, socialMedia := range socialMedias {

		var userResponse response.UserSocialMedia
		for _, user := range users {
			if user.Id == socialMedia.UserId {
				userResponse = response.UserSocialMedia{
					Id:       user.Id,
					Username: user.Username,
				}
			}

		}

		socialMediaResponse := response.GetSocialMedia{
			Id:             socialMedia.Id,
			Name:           socialMedia.Name,
			SocialMediaUrl: socialMedia.SocialMediaUrl,
			UserId:         socialMedia.UserId,
			CreatedAt:      time.Unix(int64(socialMedia.CreatedAt), 0),
			UpdatedAt:      time.Unix(int64(socialMedia.UpdatedAt), 0),
			User:           userResponse,
		}
		socialMediasResponse = append(socialMediasResponse, socialMediaResponse)
	}

	err = tx.Commit()
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			log.Println("Failed to rollback transaction:", rollbackErr)
		}
		return nil, err
	}

	return socialMediasResponse, nil
}

func (usecase *SocialMediaUsecaseImpl) UpdateSocialMedia(c context.Context, request request.SocialMedia) (response.UpdateSocialMedia, error) {
	ctx, cancel := context.WithTimeout(c, time.Duration(usecase.Timeout)*time.Second)
	defer cancel()

	err := usecase.Validate.Struct(request)
	if err != nil {
		return response.UpdateSocialMedia{}, err
	}

	tx, err := usecase.DB.Begin()
	if err != nil {
		return response.UpdateSocialMedia{}, err
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

	socialMedia := domain.SocialMedia{
		Id:             request.Id,
		Name:           request.Name,
		SocialMediaUrl: request.SocialMediaUrl,
	}

	socialMedia, err = usecase.Repository.UpdateSocialMedia(ctx, tx, socialMedia)
	if err != nil {
		return response.UpdateSocialMedia{}, err
	}

	socialMediaResponse := response.UpdateSocialMedia{
		Id:             socialMedia.Id,
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
		UserId:         socialMedia.UserId,
		UpdatedAt:      time.Unix(int64(socialMedia.UpdatedAt), 0),
	}

	err = tx.Commit()
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			log.Println("Failed to rollback transaction:", rollbackErr)
		}
		return response.UpdateSocialMedia{}, err
	}

	return socialMediaResponse, nil
}

func (usecase *SocialMediaUsecaseImpl) DeleteSocialMedia(c context.Context, id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(c, time.Duration(usecase.Timeout)*time.Second)
	defer cancel()

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

	err = usecase.Repository.DeleteSocialMedia(ctx, tx, id)
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
