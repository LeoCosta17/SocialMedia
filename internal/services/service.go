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
