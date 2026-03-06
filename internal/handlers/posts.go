package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/LeoCosta17/SocialMedia/internal/customError"
	"github.com/LeoCosta17/SocialMedia/internal/models"
	"github.com/LeoCosta17/SocialMedia/internal/request"
	"github.com/LeoCosta17/SocialMedia/internal/responses"
	"github.com/LeoCosta17/SocialMedia/internal/services"
)

type PostsHandler struct {
	services services.Service
}

func (h *PostsHandler) Create(w http.ResponseWriter, r *http.Request) {

	var post models.Post

	if err := request.ReadJSON(w, r, &post); err != nil {
		responses.BadRequestError(w, r, err)
		return
	}

	ctx := r.Context()

	post.UserID = 1

	if err := h.services.Posts.Create(ctx, &post); err != nil {
		responses.InternalServerError(w, r, err)
		return
	}

	responses.WriteJSON(w, http.StatusCreated, post)
}

func (h *PostsHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	postId, err := strconv.ParseUint(r.PathValue("post_id"), 10, 64)
	if err != nil {
		responses.BadRequestError(w, r, err)
		return
	}

	ctx := r.Context()

	post, err := h.services.Posts.GetPost(ctx, postId)
	if err != nil {
		if errors.Is(err, customError.ErrorNotFound) {
			responses.NotFoundError(w, r, err)
			return
		}
		responses.InternalServerError(w, r, err)
		return
	}
	responses.WriteJSON(w, http.StatusOK, post)
}

func (h *PostsHandler) Update(w http.ResponseWriter, r *http.Request) {

	postId, err := strconv.ParseUint(r.PathValue("post_id"), 10, 64)
	if err != nil {
		responses.BadRequestError(w, r, err)
		return
	}

	var post models.Post

	if err := request.ReadJSON(w, r, &post); err != nil {
		responses.BadRequestError(w, r, err)
		return
	}

	ctx := r.Context()

	rowsAffected, err := h.services.Posts.Update(ctx, postId, &post)
	if err != nil {
		responses.BadRequestError(w, r, err)
		return
	}

	data := map[string]uint64{
		"rowsAffected": rowsAffected,
	}

	responses.WriteJSON(w, http.StatusOK, data)
}

func (h *PostsHandler) Delete(w http.ResponseWriter, r *http.Request) {

	postId, err := strconv.ParseUint(r.PathValue("post_id"), 10, 64)
	if err != nil {
		responses.BadRequestError(w, r, err)
		return
	}

	ctx := r.Context()

	rowsAffected, err := h.services.Posts.Delete(ctx, postId)
	if err != nil {
		responses.InternalServerError(w, r, err)
		return
	}

	data := map[string]uint64{
		"rowsAffected": rowsAffected,
	}

	responses.WriteJSON(w, http.StatusOK, data)
}
