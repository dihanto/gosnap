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

type SocialMediaUsecaseImpl struct {
	Repository repository.SocialMediaRepository
	DB         *sql.DB
	Validate   *validator.Validate
}

func NewSocialMediaUsecase(repository repository.SocialMediaRepository, db *sql.DB, validate *validator.Validate) SocialMediaUsecase {
	return &SocialMediaUsecaseImpl{
		Repository: repository,
		DB:         db,
		Validate:   validate,
	}
}

func (usecase *SocialMediaUsecaseImpl) PostSocialMedia(ctx context.Context, request web.PostSocialMedia) (web.PostSocialMedia, error) {
	err := usecase.Validate.Struct(request)
	if err != nil {
		panic(err)
	}
	tx, err := usecase.DB.Begin()
	if err != nil {
		panic(err)
	}

	socialMedia := domain.SocialMedia{
		Name:           request.Name,
		SocialMediaUrl: request.SocialMediaUrl,
		UserId:         request.UserId,
	}

	socialMedia, err = usecase.Repository.PostSocialMedia(ctx, tx, socialMedia)
	if err != nil {
		panic(err)
	}

	socialMediaResponse := web.PostSocialMedia{
		Id:             socialMedia.Id,
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
		UserId:         socialMedia.UserId,
		CreatedAt:      time.Unix(int64(socialMedia.CreatedAt), 0),
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

	return socialMediaResponse, nil

}

func (usecase *SocialMediaUsecaseImpl) GetSocialMedia(ctx context.Context) ([]web.GetSocialMedia, error) {
	tx, err := usecase.DB.Begin()
	if err != nil {
		panic(err)
	}

	socialMedias, user, err := usecase.Repository.GetSocialMedia(ctx, tx)
	if err != nil {
		panic(err)
	}
	userWeb := web.UserSocialMedia{
		Id:       user.Id,
		Username: user.Username,
	}

	var socialMediasResponse []web.GetSocialMedia
	for _, socialMedia := range socialMedias {
		socialMedia := web.GetSocialMedia{
			Id:             socialMedia.Id,
			Name:           socialMedia.Name,
			SocialMediaUrl: socialMedia.SocialMediaUrl,
			UserId:         socialMedia.UserId,
			CreatedAt:      time.Unix(int64(socialMedia.CreatedAt), 0),
			UpdatedAt:      time.Unix(int64(socialMedia.UpdatedAt), 0),
			User:           userWeb,
		}
		socialMediasResponse = append(socialMediasResponse, socialMedia)
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

	return socialMediasResponse, nil
}

func (usecase *SocialMediaUsecaseImpl) UpdateSocialMedia(ctx context.Context, request web.UpdateSocialMedia) (web.UpdateSocialMedia, error) {
	err := usecase.Validate.Struct(request)
	if err != nil {
		panic(err)
	}
	tx, err := usecase.DB.Begin()
	if err != nil {
		panic(err)
	}

	socialMedia := domain.SocialMedia{
		Id:             request.Id,
		Name:           request.Name,
		SocialMediaUrl: request.SocialMediaUrl,
	}

	socialMedia, err = usecase.Repository.UpdateSocialMedia(ctx, tx, socialMedia)
	if err != nil {
		panic(err)
	}

	socialMediaResponse := web.UpdateSocialMedia{
		Id:             socialMedia.Id,
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
		UserId:         socialMedia.UserId,
		UpdatedAt:      time.Unix(int64(socialMedia.UpdatedAt), 0),
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

	return socialMediaResponse, nil
}

func (usecase *SocialMediaUsecaseImpl) DeleteSocialMedia(ctx context.Context, id int) error {
	tx, err := usecase.DB.Begin()
	if err != nil {
		panic(err)
	}

	err = usecase.Repository.DeleteSocialMedia(ctx, tx, id)
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
