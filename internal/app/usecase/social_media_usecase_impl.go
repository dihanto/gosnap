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

type SocialMediaUsecaseImpl struct {
	Repository repository.SocialMediaRepository
	Validate   *validator.Validate
	Timeout    int
}

func NewSocialMediaUsecase(repository repository.SocialMediaRepository, validate *validator.Validate, timeout int) SocialMediaUsecase {
	return &SocialMediaUsecaseImpl{
		Repository: repository,
		Validate:   validate,
		Timeout:    timeout,
	}
}

func (usecase *SocialMediaUsecaseImpl) PostSocialMedia(ctx context.Context, request request.SocialMedia) (response.PostSocialMedia, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(usecase.Timeout)*time.Second)
	defer cancel()

	err := usecase.Validate.Struct(request)
	if err != nil {
		return response.PostSocialMedia{}, err
	}

	socialMedia := domain.SocialMedia{
		Name:           request.Name,
		SocialMediaUrl: request.SocialMediaUrl,
		UserId:         request.UserId,
	}

	socialMedia, err = usecase.Repository.PostSocialMedia(ctx, socialMedia)
	if err != nil {
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

func (usecase *SocialMediaUsecaseImpl) GetSocialMedia(ctx context.Context) ([]response.GetSocialMedia, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(usecase.Timeout)*time.Second)
	defer cancel()

	socialMedias, users, err := usecase.Repository.GetSocialMedia(ctx)
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

	return socialMediasResponse, nil
}

func (usecase *SocialMediaUsecaseImpl) UpdateSocialMedia(ctx context.Context, request request.SocialMedia) (response.UpdateSocialMedia, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(usecase.Timeout)*time.Second)
	defer cancel()

	err := usecase.Validate.Struct(request)
	if err != nil {
		return response.UpdateSocialMedia{}, err
	}

	socialMedia := domain.SocialMedia{
		Id:             request.Id,
		Name:           request.Name,
		SocialMediaUrl: request.SocialMediaUrl,
	}

	socialMedia, err = usecase.Repository.UpdateSocialMedia(ctx, socialMedia)
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

	return socialMediaResponse, nil
}

func (usecase *SocialMediaUsecaseImpl) DeleteSocialMedia(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(usecase.Timeout)*time.Second)
	defer cancel()

	err := usecase.Repository.DeleteSocialMedia(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
