package handlers

import (
	"encoding/json"
	"github.com/HeadGardener/blog-app/api-service/internal/app/models"
	"github.com/HeadGardener/blog-app/api-service/pkg/responses"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (h *Handler) createPost(w http.ResponseWriter, r *http.Request) {
	userID, err := GetUserID(r)
	if err != nil {
		responses.NewErrResponse(w, http.StatusBadRequest, err.Error(), h.errLogger)
		return
	}

	var postInput models.CreatePostInput

	if err := json.NewDecoder(r.Body).Decode(&postInput); err != nil {
		responses.NewErrResponse(w, http.StatusBadRequest, "invalid data to decode post input", h.errLogger)
		return
	}

	if err := postInput.Validate(); err != nil {
		responses.NewErrResponse(w, http.StatusBadRequest, err.Error(), h.errLogger)
		return
	}

	postInput.UserID = userID

	postID, err := h.service.CreatePost(r.Context(), postInput)
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

	post, err := h.service.GetPostByID(r.Context(), postID)
	if err != nil {
		responses.NewErrResponse(w, http.StatusInternalServerError, err.Error(), h.errLogger)
		return
	}

	responses.NewResponse(w, http.StatusOK, post)
}

func (h *Handler) getUserPosts(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	postsAmount := r.URL.Query().Get("amount")

	_, err := strconv.Atoi(postsAmount)
	if err != nil {
		postsAmount = "0"
	}

	posts, err := h.service.GetPosts(r.Context(), userID, postsAmount)
	if err != nil {
		responses.NewErrResponse(w, http.StatusInternalServerError, err.Error(), h.errLogger)
		return
	}

	responses.NewResponse(w, http.StatusOK, posts)
}

func (h *Handler) updatePost(w http.ResponseWriter, r *http.Request) {
	var postInput models.UpdatePostInput
	if err := json.NewDecoder(r.Body).Decode(&postInput); err != nil {
		responses.NewErrResponse(w, http.StatusBadRequest, "invalid data to decode post input", h.errLogger)
		return
	}

	userID, err := GetUserID(r)
	if err != nil {
		responses.NewErrResponse(w, http.StatusBadRequest, err.Error(), h.errLogger)
		return
	}
	postInput.UserID = userID

	postID := chi.URLParam(r, "post_id")
	postInput.ID = postID

	post, err := h.service.UpdatePost(r.Context(), postInput)
	if err != nil {
		responses.NewErrResponse(w, http.StatusInternalServerError, err.Error(), h.errLogger)
	}

	responses.NewResponse(w, http.StatusOK, post)
}

func (h *Handler) deletePost(w http.ResponseWriter, r *http.Request) {
	userID, err := GetUserID(r)
	if err != nil {
		responses.NewErrResponse(w, http.StatusBadRequest, err.Error(), h.errLogger)
		return
	}

	postID := chi.URLParam(r, "post_id")

	post, err := h.service.DeletePost(r.Context(), postID, userID)
	if err != nil {
		responses.NewErrResponse(w, http.StatusInternalServerError, err.Error(), h.errLogger)
	}

	responses.NewResponse(w, http.StatusOK, post)
}
