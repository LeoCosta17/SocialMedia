package models

import "errors"

type User struct {
	ID        uint64 `json:"id,omitempty"`
	Username  string `json:"username,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

func (u *User) ValidateOnCreationUpdate() error {
	if u.Username == "" {
		return errors.New("username required")
	}
	if u.Email == "" {
		return errors.New("email required")
	}
	if u.Password == "" {
		return errors.New("password required")
	}
	return nil
}
