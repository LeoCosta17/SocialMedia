package handlers

import (
	"net/http"

	"github.com/LeoCosta17/SocialMedia/internal/services"
)

type FeedHandler struct {
	services services.Service
}

func (h *FeedHandler) GetFeed(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	posts, err := h.services.Posts
}
