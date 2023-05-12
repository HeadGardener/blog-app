package handlers

import (
	"encoding/json"
	"github.com/HeadGardener/blog-app/comment-service/internal/app/models"
	"net/http"
)

func (h *Handler) createComment(w http.ResponseWriter, r *http.Request) {
	var commentInput models.CreateCommentInput

	if err := json.NewDecoder(r.Body).Decode(&commentInput); err != nil {
		h.newErrResponse(w, http.StatusBadRequest, "invalid data to decode comment input")
		return
	}

	if err := commentInput.Validate(); err != nil {
		h.newErrResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	commentID, err := h.service.CommentInterface.Create(r.Context(), commentInput)
	if err != nil {
		h.newErrResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	newResponse(w, http.StatusCreated, map[string]interface{}{
		"id": commentID,
	})
}
