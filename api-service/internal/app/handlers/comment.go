package handlers

import (
	"encoding/json"
	"github.com/HeadGardener/blog-app/api-service/internal/app/models"
	"github.com/HeadGardener/blog-app/api-service/pkg/responses"
	"net/http"
)

func (h *Handler) createComment(w http.ResponseWriter, r *http.Request) {
	userID, err := GetUserID(r)
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
