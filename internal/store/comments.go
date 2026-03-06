package store

import (
	"context"
	"database/sql"

	"github.com/LeoCosta17/SocialMedia/internal/models"
)

type CommentsStore struct {
	db *sql.DB
}

func (s *CommentsStore) Create(ctx context.Context, comment *models.Comment) error {

	query := `
		INSERT INTO comments (post_id, user_id, content, created_at)
		VALUES ($1, $2, $3, CURRENT_TIMESTAMP)
		RETURNING id, created_at
	`

	if err := s.db.QueryRowContext(
		ctx,
		query,
		comment.Post_id,
		comment.User_id,
		comment.Content,
	).Scan(
		&comment.ID,
		&comment.Created_at,
	); err != nil {
		return err
	}

	return nil
}

func (s *CommentsStore) GetByPostId(ctx context.Context, postID uint64) ([]models.Comment, error) {

	query := `
		SELECT c.id, c.post_id, c.user_id, c.content, c.created_at, u.username
		FROM comments c JOIN users u ON u.id = c.user_id
		WHERE c.post_id = $1
		ORDER BY c.created_at DESC;
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeOut)
	defer cancel()

	rows, err := s.db.QueryContext(
		ctx,
		query,
		postID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := []models.Comment{}

	for rows.Next() {

		var comment models.Comment

		if err := rows.Scan(
			&comment.ID,
			&comment.Post_id,
			&comment.User_id,
			&comment.Content,
			&comment.Created_at,
			&comment.User.Username,
		); err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	return comments, nil

}
