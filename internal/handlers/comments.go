package handlers

import (
	"net/http"
	"strconv"

	"github.com/LeoCosta17/SocialMedia/internal/models"
	"github.com/LeoCosta17/SocialMedia/internal/request"
	"github.com/LeoCosta17/SocialMedia/internal/responses"
	"github.com/LeoCosta17/SocialMedia/internal/services"
)

type CommentsHandler struct {
	services services.Service
}

/*
TODO: Validate if the comment`s post really exists
*/
func (h *CommentsHandler) Create(w http.ResponseWriter, r *http.Request) {
	userId := 1

	postId, err := strconv.ParseUint(r.PathValue("post_id"), 10, 64)
	if err != nil {
		responses.BadRequestError(w, r, err)
		return
	}

	var comment models.Comment

	if err := request.ReadJSON(w, r, &comment); err != nil {
		responses.BadRequestError(w, r, err)
		return
	}

	comment.Post_id = postId
	comment.User_id = uint64(userId)

	ctx := r.Context()

	if err := h.services.Comments.Create(ctx, &comment); err != nil {
		responses.InternalServerError(w, r, err)
		return
	}

	responses.WriteJSON(w, http.StatusCreated, comment)
}

func (h *CommentsHandler) GetAll(w http.ResponseWriter, r *http.Request) {

	postID, err := strconv.ParseUint(r.PathValue("post_id"), 10, 64)
	if err != nil {
		responses.BadRequestError(w, r, err)
		return
	}

	ctx := r.Context()

	comments, err := h.services.Comments.GetByPostId(ctx, postID)
	if err != nil {
		responses.InternalServerError(w, r, err)
		return
	}

	responses.WriteJSON(w, http.StatusOK, comments)
}
