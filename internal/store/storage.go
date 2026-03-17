package store

import (
	"context"
	"database/sql"
	"time"

	"github.com/LeoCosta17/SocialMedia/internal/models"
)

// Defines some useful variables for reuse in this package
var (
	QueryTimeOut  = time.Second * 5
	InsertTimeOut = time.Second * 10
)

// Create the storage structure
type Storage struct {
	Posts interface {
		Create(context.Context, *models.Post) error
		GetPost(context.Context, uint64) (*models.Post, error)
		GetUserFeed(ctx context.Context, userID uint64, feedQuery models.PaginatedFeedQuery) ([]models.PostFeed, error)
		Update(ctx context.Context, postId uint64, post *models.Post) (uint64, error)
		Delete(ctx context.Context, postId uint64) (uint64, error)
	}
	Users interface {
		Create(context.Context, *models.User) error
		GetByID(context.Context, uint64) (*models.User, error)
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

// Create a new instance of storage
func NewPostgresStorage(db *sql.DB) Storage {
	return Storage{
		Posts:    &PostsStorage{db: db},
		Users:    &UsersStorage{db: db},
		Comments: &CommentsStore{db: db},
	}
}
