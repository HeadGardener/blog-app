package auth

import (
	"encoding/json"
	"github.com/HeadGardener/blog-app/api-service/internal/app/models"
	jwt_helper "github.com/HeadGardener/blog-app/api-service/pkg/jwt-helper"
	"github.com/HeadGardener/blog-app/api-service/pkg/responses"
	"net/http"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	var userInput models.CreateUserInput

	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		responses.NewErrResponse(w, http.StatusBadRequest, err.Error(), h.errLogger)
		return
	}

	if err := userInput.Validate(); err != nil {
		responses.NewErrResponse(w, http.StatusBadRequest, err.Error(), h.errLogger)
		return
	}

	userID, err := h.userService.Create(r.Context(), userInput)
	if err != nil {
		responses.NewErrResponse(w, http.StatusInternalServerError, err.Error(), h.errLogger)
		return
	}

	responses.NewResponse(w, http.StatusCreated, map[string]interface{}{
		"id": userID,
	})
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	var userInput models.LogUserInput

	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		responses.NewErrResponse(w, http.StatusBadRequest, err.Error(), h.errLogger)
		return
	}

	if err := userInput.Validate(); err != nil {
		responses.NewErrResponse(w, http.StatusBadRequest, err.Error(), h.errLogger)
		return
	}

	user, err := h.userService.GetUser(r.Context(), userInput)
	if err != nil {
		responses.NewErrResponse(w, http.StatusInternalServerError, err.Error(), h.errLogger)
		return
	}

	token, err := jwt_helper.GenerateToken(user)
	if err != nil {
		responses.NewErrResponse(w, http.StatusInternalServerError, err.Error(), h.errLogger)
		return
	}

	responses.NewResponse(w, http.StatusCreated, map[string]interface{}{
		"token": token,
	})
}
