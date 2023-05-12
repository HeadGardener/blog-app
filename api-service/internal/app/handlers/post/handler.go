package post

import (
	"github.com/HeadGardener/blog-app/api-service/internal/app/handlers"
	"github.com/HeadGardener/blog-app/api-service/internal/app/handlers/middleware"
	postService "github.com/HeadGardener/blog-app/api-service/internal/app/services/post"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type Handler struct {
	postService postService.PostService
	errLogger   *zap.Logger
}

func NewPostHandler(service postService.PostService) *Handler {
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
		r.Get("/{post_id}", h.getByID)
		r.Get("/", h.getUserPosts)
		r.Put("/{post_id}", h.updatePost)
		r.Delete("/{post_id}", h.deletePost)
	})

	router.Mount("/api/post", r)
}
