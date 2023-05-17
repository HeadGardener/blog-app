package handlers

import (
	"encoding/json"
	"github.com/HeadGardener/blog-app/api-service/internal/app/models"
	"github.com/HeadGardener/blog-app/api-service/pkg/responses"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (h *Handler) createComment(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		responses.NewErrResponse(w, http.StatusBadRequest, err.Error(), h.errLogger)
		return
	}

	var commentInput models.CreateCommentInput

	if err := json.NewDecoder(r.Body).Decode(&commentInput); err != nil {
		responses.NewErrResponse(w, http.StatusBadRequest, "invalid data to decode comment input", h.errLogger)
		return
	}

	if err := commentInput.Validate(); err != nil {
		responses.NewErrResponse(w, http.StatusBadRequest, err.Error(), h.errLogger)
		return
	}

	commentInput.UserID = userID

	commentID, err := h.service.CommentService.CreateComment(r.Context(), commentInput)
	if err != nil {
		responses.NewErrResponse(w, http.StatusInternalServerError, err.Error(), h.errLogger)
		return
	}

	responses.NewResponse(w, http.StatusCreated, map[string]interface{}{
		"id": commentID,
	})
}

func (h *Handler) getPostComments(w http.ResponseWriter, r *http.Request) {
	postID := r.URL.Query().Get("post_id")
	commentsAmount := r.URL.Query().Get("amount")

	_, err := strconv.Atoi(commentsAmount)
	if err != nil {
		commentsAmount = "0"
	}

	comments, err := h.service.CommentService.GetComments(r.Context(), postID, commentsAmount)
	if err != nil {
		responses.NewErrResponse(w, http.StatusInternalServerError, err.Error(), h.errLogger)
		return
	}

	responses.NewResponse(w, http.StatusOK, comments)
}

func (h *Handler) updateComment(w http.ResponseWriter, r *http.Request) {
	var commentInput models.UpdateCommentInput

	if err := json.NewDecoder(r.Body).Decode(&commentInput); err != nil {
		responses.NewErrResponse(w, http.StatusBadRequest, "invalid data to decode comment input", h.errLogger)
		return
	}

	userID, err := getUserID(r)
	if err != nil {
		responses.NewErrResponse(w, http.StatusBadRequest, err.Error(), h.errLogger)
		return
	}
	commentInput.UserID = userID

	commentInput.ID = chi.URLParam(r, "comment_id")

	comment, err := h.service.CommentService.UpdateComment(r.Context(), commentInput)
	if err != nil {
		responses.NewErrResponse(w, http.StatusInternalServerError, err.Error(), h.errLogger)
		return
	}

	responses.NewResponse(w, http.StatusOK, comment)
}

func (h *Handler) deleteComment(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		responses.NewErrResponse(w, http.StatusBadRequest, err.Error(), h.errLogger)
		return
	}

	commentID := chi.URLParam(r, "comment_id")

	comment, err := h.service.CommentService.DeleteComment(r.Context(), commentID, userID)
	if err != nil {
		responses.NewErrResponse(w, http.StatusInternalServerError, err.Error(), h.errLogger)
		return
	}

	responses.NewResponse(w, http.StatusOK, comment)
}
