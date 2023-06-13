package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/dihanto/gosnap/model/domain"
	"github.com/google/uuid"
)

type SocialMediaRepositoryImpl struct {
}

func NewSocialMediaRepository() SocialMediaRepository {
	return &SocialMediaRepositoryImpl{}
}
func (repository *SocialMediaRepositoryImpl) PostSocialMedia(ctx context.Context, tx *sql.Tx, socialMedia domain.SocialMedia) (domain.SocialMedia, error) {
	socialMedia.CreatedAt = int32(time.Now().Unix())

	query := "INSERT INTO social_medias (id ,name, social_media_url, user_id, created_at) VALUES ($1, $2, $3, $4, $5)"
	_, err := tx.ExecContext(ctx, query, socialMedia.Id, socialMedia.Name, socialMedia.SocialMediaUrl, socialMedia.UserId, socialMedia.CreatedAt)
	if err != nil {
		return domain.SocialMedia{}, err
	}

	return socialMedia, nil

}

func (repository *SocialMediaRepositoryImpl) GetSocialMedia(ctx context.Context, tx *sql.Tx) ([]domain.SocialMedia, []domain.User, error) {
	query := "SELECT social_medias.id, social_medias.name, social_medias.social_media_url, social_medias.user_id, social_medias.created_at, social_medias.updated_at, users.id, users.username FROM social_medias JOIN users ON social_medias.user_id = users.id WHERE social_medias.deleted_at IS NULL;"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []domain.SocialMedia{}, []domain.User{}, err
	}
	defer rows.Close()

	socialMedias := []domain.SocialMedia{}
	users := []domain.User{}
	for rows.Next() {
		socialMedia := domain.SocialMedia{}
		user := domain.User{}
		err = rows.Scan(&socialMedia.Id, &socialMedia.Name, &socialMedia.SocialMediaUrl, &socialMedia.UserId, &socialMedia.CreatedAt, &socialMedia.UpdatedAt, &user.Id, &user.Username)
		if err != nil {
			return []domain.SocialMedia{}, []domain.User{}, err
		}
		user.Id = socialMedia.UserId
		users = append(users, user)
		socialMedias = append(socialMedias, socialMedia)
	}

	return socialMedias, users, nil
}

func (repository *SocialMediaRepositoryImpl) UpdateSocialMedia(ctx context.Context, tx *sql.Tx, socialMedia domain.SocialMedia) (domain.SocialMedia, error) {
	socialMedia.UpdatedAt = int32(time.Now().Unix())

	query := "UPDATE social_medias SET name=$1, social_media_url=$2, updated_at=$3 WHERE id=$4 RETURNING user_id"
	row := tx.QueryRowContext(ctx, query, socialMedia.Name, socialMedia.SocialMediaUrl, socialMedia.UpdatedAt, socialMedia.Id)

	err := row.Scan(&socialMedia.UserId)
	if err != nil {
		return domain.SocialMedia{}, err
	}

	return socialMedia, nil

}

func (repository *SocialMediaRepositoryImpl) DeleteSocialMedia(ctx context.Context, tx *sql.Tx, id uuid.UUID) error {
	deleteTime := int32(time.Now().Unix())

	query := "UPDATE social_medias SET deleted_at=$1 WHERE id=$2"
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
