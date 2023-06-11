package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/dihanto/gosnap/model/domain"
)

type SocialMediaRepositoryImpl struct {
}

func NewSocialMediaRepository() SocialMediaRepository {
	return &SocialMediaRepositoryImpl{}
}
func (repository *SocialMediaRepositoryImpl) PostSocialMedia(ctx context.Context, tx *sql.Tx, socialMedia domain.SocialMedia) (domain.SocialMedia, error) {
	t := time.Now()
	socialMedia.CreatedAt = int32(t.Unix())

	query := "insert into social_medias (name, social_media_url, user_id, created_at) values ($1, $2, $3, $4) returning id"
	row := tx.QueryRowContext(ctx, query, socialMedia.Name, socialMedia.SocialMediaUrl, socialMedia.UserId, socialMedia.CreatedAt)

	err := row.Scan(&socialMedia.Id)
	if err != nil {
		return domain.SocialMedia{}, err
	}

	return socialMedia, nil

}

func (repository *SocialMediaRepositoryImpl) GetSocialMedia(ctx context.Context, tx *sql.Tx) ([]domain.SocialMedia, domain.User, error) {
	query := "select social_medias.id, social_medias.name, social_medias.social_media_url, social_medias.user_id, social_medias.created_at, social_medias.updated_at, users.id, users.username from social_medias join users on social_medias.user_id = users.id where social_medias.deleted_at is null;"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []domain.SocialMedia{}, domain.User{}, err
	}
	defer rows.Close()

	socialMedias := []domain.SocialMedia{}
	user := domain.User{}
	for rows.Next() {
		socialMedia := domain.SocialMedia{}
		err = rows.Scan(&socialMedia.Id, &socialMedia.Name, &socialMedia.SocialMediaUrl, &socialMedia.UserId, &socialMedia.CreatedAt, &socialMedia.UpdatedAt, &user.Id, &user.Username)
		if err != nil {
			return []domain.SocialMedia{}, domain.User{}, err
		}
		socialMedias = append(socialMedias, socialMedia)
	}

	return socialMedias, user, nil
}

func (repository *SocialMediaRepositoryImpl) UpdateSocialMedia(ctx context.Context, tx *sql.Tx, socialMedia domain.SocialMedia) (domain.SocialMedia, error) {
	t := time.Now()
	socialMedia.UpdatedAt = int32(t.Unix())

	query := "update social_medias set name=$1, social_media_url=$2, updated_at=$3 where id=$4 returning user_id"
	row := tx.QueryRowContext(ctx, query, socialMedia.Name, socialMedia.SocialMediaUrl, socialMedia.UpdatedAt, socialMedia.Id)

	err := row.Scan(&socialMedia.UserId)
	if err != nil {
		return domain.SocialMedia{}, err
	}

	return socialMedia, nil

}

func (repository *SocialMediaRepositoryImpl) DeleteSocialMedia(ctx context.Context, tx *sql.Tx, id int) error {
	t := time.Now()
	deleteTime := int32(t.Unix())

	query := "update social_medias set deleted_at=$1 where id=$2"
	result, err := tx.ExecContext(ctx, query, deleteTime, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("social media not found")
	}

	return nil
}
