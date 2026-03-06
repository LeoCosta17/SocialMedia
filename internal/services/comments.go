package services

import (
	"context"

	"github.com/LeoCosta17/SocialMedia/internal/models"
	"github.com/LeoCosta17/SocialMedia/internal/store"
)

type CommentsService struct {
	store store.Storage
}

/*
* * Comment Struct data/fields:
	ID         uint64 `json:"id,omitempty"`
	Post_id    uint64 `json:"post_id,omitempty"`
	User_id    uint64 `json:"user_id,omitempty"`
	Content    string `json:"content,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	User       User   `json:"user,omitempty"`

*/

func (s *CommentsService) Create(ctx context.Context, comment *models.Comment) error {

	if err := comment.ValidateOnCreationUpdate(); err != nil {
		return err
	}

	if err := s.store.Comments.Create(ctx, comment); err != nil {
		return err
	}

	return nil

}

func (s *CommentsService) GetByPostId(ctx context.Context, postId uint64) ([]models.Comment, error) {

	comments, err := s.store.Comments.GetByPostId(ctx, postId)
	if err != nil {
		return nil, err
	}

	return comments, nil
}
