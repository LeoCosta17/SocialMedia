package services

import (
	"context"

	"github.com/LeoCosta17/SocialMedia/internal/models"
	"github.com/LeoCosta17/SocialMedia/internal/store"
)

type Service struct {
	Posts interface {
		Create(ctx context.Context, post *models.Post) error
		GetPost(ctx context.Context, postId uint64) (*models.Post, error)
		Update(ctx context.Context, postId uint64, post *models.Post) (uint64, error)
		Delete(ctx context.Context, postId uint64) (uint64, error)
	}
	Users interface {
		Create(context.Context, *models.User) error
		GetByID(context.Context, uint64) (*models.User, error)
		// Receives the context of the request, followerID and user that will be followed ID
		// Returns the number of inserted rows and a error
		Follow(context.Context, uint64, uint64) (uint64, error)
		// Receives the request context, follower ID and the user ID
		// Should return a uint64 (and should be 1) and an error
		Unfollow(context.Context, uint64, uint64) (uint64, error)
	}
	Comments interface {
		Create(context.Context, *models.Comment) error
		GetByPostId(context.Context, uint64) ([]models.Comment, error)
	}
}

func NewService(store store.Storage) Service {
	return Service{
		Posts:    &PostsService{store: store},
		Users:    &UsersService{store: store},
		Comments: &CommentsService{store: store},
	}
}
