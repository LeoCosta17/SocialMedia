package services

import (
	"context"
	"errors"

	"github.com/LeoCosta17/SocialMedia/internal/models"
	"github.com/LeoCosta17/SocialMedia/internal/store"
)

type UsersService struct {
	store store.Storage
}

func NewUsersService(store store.Storage) UsersService {
	return UsersService{store: store}
}

func (s *UsersService) Create(ctx context.Context, user *models.User) error {

	if err := user.ValidateOnCreationUpdate(); err != nil {
		return err
	}

	if err := s.store.Users.Create(ctx, user); err != nil {
		return err
	}

	return nil
}

func (s *UsersService) GetByID(ctx context.Context, userId uint64) (*models.User, error) {

	user, err := s.store.Users.GetByID(ctx, userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UsersService) Follow(ctx context.Context, followerID, followedID uint64) (uint64, error) {

	if followerID == followedID {
		return 0, errors.New("you cannot follow yourself!")
	}

	rowsInserted, err := s.store.Users.Follow(ctx, followerID, followedID)
	if err != nil {
		return 0, err
	}

	return rowsInserted, nil

}

func (s *UsersService) Unfollow(ctx context.Context, followerID, followedID uint64) (uint64, error) {
	if followerID == followedID {
		return 0, errors.New("you cannot unfollow yourself!")
	}

	rowsInserted, err := s.store.Users.Unfollow(ctx, followerID, followedID)
	if err != nil {
		return 0, err
	}

	return rowsInserted, nil
}
