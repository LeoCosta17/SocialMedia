package store

import (
	"context"
	"database/sql"
	"errors"

	"github.com/LeoCosta17/SocialMedia/internal/customError"
	"github.com/LeoCosta17/SocialMedia/internal/models"
)

type UsersStorage struct {
	db *sql.DB
}

func (s *UsersStorage) Create(ctx context.Context, user *models.User) error {

	query := `
		INSERT INTO users(username, email, password)
		VALUES($1, $2, $3)
		RETURNING id, created_at
	`

	if err := s.db.QueryRowContext(
		ctx,
		query,
		user.Username,
		user.Email,
		user.Password,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	); err != nil {
		return err
	}

	return nil
}

func (s *UsersStorage) GetByID(ctx context.Context, userId uint64) (*models.User, error) {

	query := `
		SELECT id, username, email, created_at FROM users WHERE id = $1
	`

	var user models.User

	if err := s.db.QueryRowContext(
		ctx,
		query,
		userId,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.CreatedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, customError.ErrorNotFound
		}
		return nil, err
	}

	return &user, nil
}
