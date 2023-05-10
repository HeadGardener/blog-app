package auth

import (
	"github.com/HeadGardener/blog-app/api-service/internal/app/handlers"
	userService "github.com/HeadGardener/blog-app/api-service/internal/app/services/user"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type Handler struct {
	userService userService.UserService
	errLogger   *zap.Logger
}

func NewAuthHandler(service userService.UserService) *Handler {
	return &Handler{
		userService: service,
		errLogger:   handlers.NewLogger(),
	}
}

func (h *Handler) InitRoutes(router *chi.Mux) {
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Post("/sign-up", h.signUp)
		r.Post("/sign-in", h.signIn)
	})

	router.Mount("/api/auth", r)
}
