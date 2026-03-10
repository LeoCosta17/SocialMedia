package models

import (
	"errors"
)

type Post struct {
	ID        uint64    `json:"id,omitempty"`
	Content   string    `json:"content,omitempty"`
	Title     string    `json:"title,omitempty"`
	UserID    uint64    `json:"user_id,omitempty"`
	Tags      []string  `json:"tags,omitempty"`
	Comments  []Comment `json:"comments,omitempty"`
	User      User      `json:"user,omitempty"`
	CreatedAt string    `json:"created_at,omitempty"`
	UpdatedAt string    `json:"updated_at,omitempty"`
}

func (p *Post) ValidateOnCreationUpdate() error {
	if p.Content == "" {
		return errors.New("content required")
	}
	if p.Title == "" {
		return errors.New("title required")
	}
	if len(p.Tags) == 0 {
		return errors.New("tags required")
	}
	return nil
}
