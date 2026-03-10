package store

import (
	"context"
	"database/sql"
	"errors"

	"github.com/LeoCosta17/SocialMedia/internal/customError"
	"github.com/LeoCosta17/SocialMedia/internal/models"
	"github.com/lib/pq"
)

type PostsStorage struct {
	db *sql.DB
}

func (s *PostsStorage) Create(ctx context.Context, post *models.Post) error {

	query := `
		INSERT INTO posts (content, title, user_id, tags)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeOut)
	defer cancel()

	if err := s.db.QueryRowContext(
		ctx,
		query,
		post.Content,
		post.Title,
		post.UserID,
		pq.Array(&post.Tags),
	).Scan(
		&post.ID,
		&post.CreatedAt,
	); err != nil {
		return err
	}

	return nil
}

func (s *PostsStorage) GetPost(ctx context.Context, postId uint64) (*models.Post, error) {

	query := `
		SELECT id, title, user_id, content, tags, created_at, updated_at
		FROM posts
		WHERE id = $1
	`

	post := &models.Post{}

	ctx, cancel := context.WithTimeout(ctx, QueryTimeOut)
	defer cancel()

	if err := s.db.QueryRowContext(
		ctx,
		query,
		postId,
	).Scan(
		&post.ID,
		&post.Title,
		&post.UserID,
		&post.Content,
		pq.Array(&post.Tags),
		&post.CreatedAt,
		&post.UpdatedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, customError.ErrorNotFound
		}
		return nil, err
	}

	query = `
		SELECT c.id, c.content, c.created_at, u.username, u.email
		FROM comments c INNER JOIN users u ON c.user_id = u.id
		WHERE c.post_id = $1
	`

	rows, err := s.db.QueryContext(ctx, query, postId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {

		var comment models.Comment

		if err := rows.Scan(
			&comment.ID,
			&comment.Content,
			&comment.Created_at,
			&comment.User.Username,
			&comment.User.Email,
		); err != nil {
			return nil, err
		}

		post.Comments = append(post.Comments, comment)
	}

	return post, nil
}

func (s *PostsStorage) GetUserFeed(ctx context.Context, userID uint64) ([]models.Post, error) {

}

func (s *PostsStorage) Update(ctx context.Context, postId uint64, post *models.Post) (uint64, error) {

	query := `
		UPDATE posts
		SET title = $1, content = $2
		WHERE id =$3
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeOut)
	defer cancel()

	result, err := s.db.ExecContext(
		ctx,
		query,
		post.Title,
		post.Content,
		postId,
	)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return uint64(rowsAffected), nil
}
func (s *PostsStorage) Delete(ctx context.Context, postId uint64) (uint64, error) {

	query := `
		DELETE FROM posts WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeOut)
	defer cancel()

	result, err := s.db.ExecContext(ctx, query, postId)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return uint64(rowsAffected), nil

}
