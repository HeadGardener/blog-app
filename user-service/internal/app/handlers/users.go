package handlers

import (
	"encoding/json"
	"github.com/HeadGardener/user-service/internal/app/models"
	"net/http"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	var userInput models.CreateUserInput

	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		h.newErrResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := userInput.Validate(); err != nil {
		h.newErrResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	userID, err := h.service.Create(r.Context(), userInput)
	if err != nil {
		h.newErrResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(w, http.StatusCreated, map[string]interface{}{
		"id": userID,
	})
}
