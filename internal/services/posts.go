package services

import (
	"context"

	"github.com/LeoCosta17/SocialMedia/internal/models"
	"github.com/LeoCosta17/SocialMedia/internal/store"
)

type PostsService struct {
	store store.Storage
}

func NewPostsService(store store.Storage) PostsService {
	return PostsService{store: store}
}

func (s *PostsService) Create(ctx context.Context, post *models.Post) error {

	if err := post.ValidateOnCreationUpdate(); err != nil {
		return err
	}

	if err := s.store.Posts.Create(ctx, post); err != nil {
		return err
	}

	return nil
}

func (s *PostsService) GetPost(ctx context.Context, postId uint64) (*models.Post, error) {

	post, err := s.store.Posts.GetPost(ctx, postId)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (s *PostsService) GetUserFeed(ctx context.Context, userID uint64) ([]models.Post, error) {

	feed, err := s.store.Posts.GetUserFeed(ctx, userID)
	if err != nil {
		return nil, err
	}

	return feed, nil
}

func (s *PostsService) Update(ctx context.Context, postId uint64, post *models.Post) (uint64, error) {

	if err := post.ValidateOnCreationUpdate(); err != nil {
		return 0, err
	}

	rowsAffected, err := s.store.Posts.Update(ctx, postId, post)
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (s *PostsService) Delete(ctx context.Context, postId uint64) (uint64, error) {

	rowsAffected, err := s.store.Posts.Delete(ctx, postId)
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
