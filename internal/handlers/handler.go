package handlers

import (
	"net/http"

	service "github.com/LeoCosta17/SocialMedia/internal/services"
)

type Handler struct {
	Users interface {
		// Create a new user
		Create(http.ResponseWriter, *http.Request)
		// Get a user by ID
		GetByID(http.ResponseWriter, *http.Request)
		// Follow a user
		Follow(http.ResponseWriter, *http.Request)
		// Unfollow a user
		Unfollow(http.ResponseWriter, *http.Request)
	}
	Posts interface {
		Create(http.ResponseWriter, *http.Request)
		GetByID(http.ResponseWriter, *http.Request)
		Update(http.ResponseWriter, *http.Request)
		Delete(http.ResponseWriter, *http.Request)
	}
	Comments interface {
		GetAll(http.ResponseWriter, *http.Request)
		Create(http.ResponseWriter, *http.Request)
	}
}

func NewHandler(service service.Service) Handler {
	return Handler{
		Users:    &UsersHandler{services: service},
		Posts:    &PostsHandler{services: service},
		Comments: &CommentsHandler{services: service},
	}
}
