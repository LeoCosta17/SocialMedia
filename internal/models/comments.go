package models

import (
	"errors"
)

type Comment struct {
	ID         uint64 `json:"id,omitempty"`
	Post_id    uint64 `json:"post_id,omitempty"`
	User_id    uint64 `json:"user_id,omitempty"`
	Content    string `json:"content,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	User       User   `json:"user,omitempty"`
}

func (c *Comment) ValidateOnCreationUpdate() error {
	if c.Post_id < 1 {
		return errors.New("post id required")
	}
	if c.User_id < 1 {
		return errors.New("user id required")
	}
	if c.Content == "" {
		return errors.New("comment content required")
	}
	return nil
}
