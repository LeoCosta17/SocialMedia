package services

import (
	"context"

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
