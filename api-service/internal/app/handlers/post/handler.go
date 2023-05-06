package post

import (
	"github.com/HeadGardener/blog-app/api-service/internal/app/handlers"
	"github.com/HeadGardener/blog-app/api-service/internal/app/handlers/middleware"
	post_service "github.com/HeadGardener/blog-app/api-service/internal/app/services/post"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type Handler struct {
	postService post_service.PostService
	errLogger   *zap.Logger
}

func NewPostHandler(service post_service.PostService) *Handler {
	return &Handler{
		postService: service,
		errLogger:   handlers.NewLogger(),
	}
}

func (h *Handler) InitRoutes(router *chi.Mux) {
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Use(mw.IdentifyUser)
		r.Post("/", h.createPost)
	})

	router.Mount("/api/post", r)
}