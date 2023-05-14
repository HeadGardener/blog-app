package handlers

import (
	"encoding/json"
	"github.com/HeadGardener/blog-app/comment-service/internal/app/models"
	"net/http"
	"strconv"
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

	commentID, err := h.service.CommentInterface.CreateComment(r.Context(), commentInput)
	if err != nil {
		h.newErrResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	newResponse(w, http.StatusCreated, map[string]interface{}{
		"id": commentID,
	})
}

func (h *Handler) getComments(w http.ResponseWriter, r *http.Request) {
	postID := r.URL.Query().Get("post_id")
	commentsAmount, err := strconv.Atoi(r.URL.Query().Get("amount"))
	if err != nil {
		h.newErrResponse(w, http.StatusBadRequest, "invalid amount param")
		return
	}

	comments, err := h.service.CommentInterface.GetComments(r.Context(), postID, commentsAmount)
	if err != nil {
		h.newErrResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(w, http.StatusOK, comments)
}

func (h *Handler) updateComment(w http.ResponseWriter, r *http.Request) {
	var commentInput models.UpdateCommentInput
	if err := json.NewDecoder(r.Body).Decode(&commentInput); err != nil {
		h.newErrResponse(w, http.StatusBadRequest, "invalid data to decode comment input")
		return
	}

	if err := commentInput.Validate(); err != nil {
		h.newErrResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	comment, err := h.service.CommentInterface.UpdateComment(r.Context(), commentInput)
	if err != nil {
		h.newErrResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(w, http.StatusOK, comment)
}
