package handlers

import (
	"net/http"

	"github.com/LeoCosta17/SocialMedia/internal/models"
	"github.com/LeoCosta17/SocialMedia/internal/responses"
	"github.com/LeoCosta17/SocialMedia/internal/services"
)

type FeedHandler struct {
	services services.Service
}

func (h *FeedHandler) GetFeed(w http.ResponseWriter, r *http.Request) {

	feedQuery := models.PaginatedFeedQuery{
		Limit:  20,
		Offset: 0,
		Sort:   "desc",
	}

	feedQuery, err := feedQuery.Parse(r)
	if err != nil {
		responses.BadRequestError(w, r, err)
		return
	}

	ctx := r.Context()

	posts, err := h.services.Posts.GetUserFeed(ctx, 2, feedQuery)
	if err != nil {
		responses.InternalServerError(w, r, err)
		return
	}

	responses.WriteJSON(w, http.StatusOK, posts)
}
