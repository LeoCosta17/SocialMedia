package handlers

import (
	"net/http"

	service "github.com/LeoCosta17/SocialMedia/internal/services"
)

type Handler struct {
	Users interface {
		Create(http.ResponseWriter, *http.Request)
		GetByID(http.ResponseWriter, *http.Request)
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
