package post

import (
	"encoding/json"
	mw "github.com/HeadGardener/blog-app/api-service/internal/app/handlers/middleware"
	"github.com/HeadGardener/blog-app/api-service/internal/app/models"
	"github.com/HeadGardener/blog-app/api-service/pkg/responses"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (h *Handler) createPost(w http.ResponseWriter, r *http.Request) {
	userID, err := mw.GetUserID(r)
	if err != nil {
		responses.NewErrResponse(w, http.StatusBadRequest, err.Error(), h.errLogger)
		return
	}

	var postInput models.CreatePostInput

	if err := json.NewDecoder(r.Body).Decode(&postInput); err != nil {
		responses.NewErrResponse(w, http.StatusBadRequest, err.Error(), h.errLogger)
		return
	}

	if err := postInput.Validate(); err != nil {
		responses.NewErrResponse(w, http.StatusBadRequest, err.Error(), h.errLogger)
		return
	}

	postInput.UserID = userID

	postID, err := h.postService.CreatePost(r.Context(), postInput)
	if err != nil {
		responses.NewErrResponse(w, http.StatusInternalServerError, err.Error(), h.errLogger)
		return
	}

	responses.NewResponse(w, http.StatusCreated, map[string]interface{}{
		"id": postID,
	})
}

func (h *Handler) getByID(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "post_id")

	post, err := h.postService.GetPostByID(r.Context(), postID)
	if err != nil {
		responses.NewErrResponse(w, http.StatusInternalServerError, err.Error(), h.errLogger)
		return
	}

	responses.NewResponse(w, http.StatusOK, post)
}
