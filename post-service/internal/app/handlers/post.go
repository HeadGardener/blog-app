package handlers

import (
	"encoding/json"
	"github.com/HeadGardener/blog-app/post-service/internal/app/models"
	"net/http"
)

func (h *Handler) createPost(w http.ResponseWriter, r *http.Request) {
	var postInput models.CreatePostInput

	if err := json.NewDecoder(r.Body).Decode(&postInput); err != nil {
		h.newErrResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := postInput.Validate(); err != nil {
		h.newErrResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	postID, err := h.service.PostInterface.Create(r.Context(), postInput)
	if err != nil {
		h.newErrResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(w, http.StatusCreated, map[string]interface{}{
		"id": postID,
	})
}
