package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/dihanto/gosnap/internal/app/helper"
	"github.com/dihanto/gosnap/model/domain"
)

type SocialMediaRepositoryImpl struct {
	Database *sql.DB
}

func NewSocialMediaRepository(database *sql.DB) SocialMediaRepository {
	return &SocialMediaRepositoryImpl{
		Database: database,
	}
}

// PostSocialMedia is a method to create a new social media entry in the database.
func (repository *SocialMediaRepositoryImpl) PostSocialMedia(ctx context.Context, socialMedia domain.SocialMedia) (domain.SocialMedia, error) {
	tx, err := repository.Database.Begin()
	if err != nil {
		return domain.SocialMedia{}, err
	}
	defer helper.CommitOrRollback(tx, &err)

	socialMedia.CreatedAt = int32(time.Now().Unix())

	query := "INSERT INTO social_medias (name, social_media_url, user_id, created_at) VALUES ($1, $2, $3, $4) RETURNING id"
	row := tx.QueryRowContext(ctx, query, socialMedia.Name, socialMedia.SocialMediaUrl, socialMedia.UserId, socialMedia.CreatedAt)
	err = row.Scan(&socialMedia.Id)
	if err != nil {
		return domain.SocialMedia{}, err
	}

	return socialMedia, nil
}

// GetSocialMedia is a method to retrieve all social media entries and their associated users from the database.
func (repository *SocialMediaRepositoryImpl) GetSocialMedia(ctx context.Context) ([]domain.SocialMedia, []domain.User, error) {
	tx, err := repository.Database.Begin()
	if err != nil {
		return []domain.SocialMedia{}, []domain.User{}, err
	}
	defer helper.CommitOrRollback(tx, &err)

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

// UpdateSocialMedia is a method to update a social media entry in the database.
func (repository *SocialMediaRepositoryImpl) UpdateSocialMedia(ctx context.Context, socialMedia domain.SocialMedia) (domain.SocialMedia, error) {
	tx, err := repository.Database.Begin()
	if err != nil {
		return domain.SocialMedia{}, err
	}
	defer helper.CommitOrRollback(tx, &err)

	socialMedia.UpdatedAt = int32(time.Now().Unix())

	query := "UPDATE social_medias SET name=$1, social_media_url=$2, updated_at=$3 WHERE id=$4 RETURNING user_id"
	row := tx.QueryRowContext(ctx, query, socialMedia.Name, socialMedia.SocialMediaUrl, socialMedia.UpdatedAt, socialMedia.Id)

	err = row.Scan(&socialMedia.UserId)
	if err != nil {
		return domain.SocialMedia{}, err
	}

	return socialMedia, nil
}

// DeleteSocialMedia is a method to "soft delete" a social media entry by setting the deleted_at field in the database.
func (repository *SocialMediaRepositoryImpl) DeleteSocialMedia(ctx context.Context, id int) error {
	tx, err := repository.Database.Begin()
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx, &err)

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
