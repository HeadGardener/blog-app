package auth

import (
	"encoding/json"
	"github.com/HeadGardener/api-service/internal/app/handlers"
	"github.com/HeadGardener/api-service/internal/app/models"
	user_service "github.com/HeadGardener/api-service/internal/app/services/user"
	"github.com/HeadGardener/api-service/pkg/responses"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
)

type Handler struct {
	userService user_service.UserService
	errLogger   *zap.Logger
}

func NewAuthHandler(service user_service.UserService) *Handler {
	return &Handler{
		userService: service,
		errLogger:   handlers.NewLogger(),
	}
}

func (h *Handler) InitRoutes(router *chi.Mux) {
	r := chi.NewRouter()
	r.Route("/auth", func(r chi.Router) {
		r.Post("/sign-up", h.signUp)
		r.Post("/sign-in", h.signIn)
	})

	router.Mount("/api", r)
}

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	var userInput models.CreateUserInput

	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		responses.NewErrResponse(w, http.StatusBadRequest, err.Error(), h.errLogger)
		return
	}

	/*if err := userInput.Validate(); err != nil {
		responses.NewErrResponse(w, http.StatusBadRequest, err.Error(), h.errLogger)
		return
	}*/

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

}
