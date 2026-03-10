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

type UsersHandler struct {
	services services.Service
}

func (h *UsersHandler) Create(w http.ResponseWriter, r *http.Request) {

	var user models.User

	if err := request.ReadJSON(w, r, &user); err != nil {
		responses.BadRequestError(w, r, err)
		return
	}

	ctx := r.Context()

	if err := h.services.Users.Create(ctx, &user); err != nil {
		responses.InternalServerError(w, r, err)
		return
	}

	responses.WriteJSON(w, http.StatusCreated, user)
}

func (h *UsersHandler) GetByID(w http.ResponseWriter, r *http.Request) {

	userId, err := strconv.ParseUint(r.PathValue("user_id"), 10, 64)
	if err != nil {
		responses.BadRequestError(w, r, err)
		return
	}

	ctx := r.Context()

	post, err := h.services.Users.GetByID(ctx, userId)
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

func (h *UsersHandler) Follow(w http.ResponseWriter, r *http.Request) {

	followerID := 5

	userToFollowID, err := strconv.ParseUint(r.PathValue("user_id"), 10, 64)
	if err != nil {
		responses.BadRequestError(w, r, err)
		return
	}

	ctx := r.Context()

	rowsInserted, err := h.services.Users.Follow(ctx, uint64(followerID), userToFollowID)
	if err != nil {
		responses.InternalServerError(w, r, err)
		return
	}

	responses.WriteJSON(w, http.StatusCreated, rowsInserted)
}

func (h *UsersHandler) Unfollow(w http.ResponseWriter, r *http.Request) {

	followerID := 5

	userID, err := strconv.ParseUint(r.PathValue("user_id"), 10, 64)
	if err != nil {
		responses.BadRequestError(w, r, err)
		return
	}

	ctx := r.Context()

	rowsInserted, err := h.services.Users.Unfollow(ctx, uint64(followerID), userID)
	if err != nil {
		responses.InternalServerError(w, r, err)
		return
	}

	responses.WriteJSON(w, http.StatusCreated, rowsInserted)

}
