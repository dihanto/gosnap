package repository

import (
	"context"
	"database/sql"

	"github.com/dihanto/gosnap/model/domain"
)

type SocialMediaRepository interface {
	PostSocialMedia(ctx context.Context, tx *sql.Tx, socialMedia domain.SocialMedia) (domain.SocialMedia, error)
	GetSocialMedia(ctx context.Context, tx *sql.Tx) ([]domain.SocialMedia, []domain.User, error)
	UpdateSocialMedia(ctx context.Context, tx *sql.Tx, socialMedia domain.SocialMedia) (domain.SocialMedia, error)
	DeleteSocialMedia(ctx context.Context, tx *sql.Tx, id int) error
}
