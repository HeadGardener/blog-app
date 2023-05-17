package handlers

import (
	"encoding/json"
	"github.com/HeadGardener/blog-app/post-service/internal/app/models"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (h *Handler) createPost(w http.ResponseWriter, r *http.Request) {
	var postInput models.CreatePostInput

	if err := json.NewDecoder(r.Body).Decode(&postInput); err != nil {
		h.newErrResponse(w, http.StatusBadRequest, "invalid data to decode post input")
		return
	}

	if err := postInput.Validate(); err != nil {
		h.newErrResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	postID, err := h.service.PostInterface.CreatePost(r.Context(), postInput)
	if err != nil {
		h.newErrResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(w, http.StatusCreated, map[string]interface{}{
		"id": postID,
	})
}

func (h *Handler) getByID(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "post_id")

	post, err := h.service.PostInterface.GetByID(r.Context(), postID)
	if err != nil {
		h.newErrResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(w, http.StatusOK, post)
}

func (h *Handler) getUserPosts(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	postsAmount, err := strconv.Atoi(r.URL.Query().Get("amount"))
	if err != nil {
		h.newErrResponse(w, http.StatusBadRequest, "invalid amount param")
		return
	}

	posts, err := h.service.PostInterface.GetPosts(r.Context(), userID, postsAmount)
	if err != nil {
		h.newErrResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(w, http.StatusOK, posts)
}

func (h *Handler) updatePost(w http.ResponseWriter, r *http.Request) {
	var postInput models.UpdatePostInput
	if err := json.NewDecoder(r.Body).Decode(&postInput); err != nil {
		h.newErrResponse(w, http.StatusBadRequest, "invalid data to decode post input")
		return
	}

	if err := postInput.Validate(); err != nil {
		h.newErrResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	post, err := h.service.PostInterface.UpdatePost(r.Context(), postInput)
	if err != nil {
		h.newErrResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(w, http.StatusOK, post)
}

func (h *Handler) deletePost(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "post_id")
	userID := r.URL.Query().Get("user_id")

	post, err := h.service.PostInterface.DeletePost(r.Context(), postID, userID)
	if err != nil {
		h.newErrResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(w, http.StatusOK, post)
}
