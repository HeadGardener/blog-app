package handlers

import (
	"encoding/json"
	"github.com/HeadGardener/blog-app/api-service/internal/app/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type Handler struct {
	service   *services.Service
	errLogger *zap.Logger
}

func NewHandler(service *services.Service) *Handler {
	return &Handler{
		service:   service,
		errLogger: NewLogger(),
	}
}

func (h *Handler) InitRoutes() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/api", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Post("/sign-up", h.signUp)
			r.Post("/sign-in", h.signIn)
		})

		r.Route("/post", func(r chi.Router) {
			r.Use(IdentifyUser)
			r.Post("/", h.createPost)
			r.Get("/{post_id}", h.getByID)
			r.Get("/", h.getUserPosts)
			r.Put("/{post_id}", h.updatePost)
			r.Delete("/{post_id}", h.deletePost)
		})

		r.Route("/comment", func(r chi.Router) {
			r.Use(IdentifyUser)
			r.Post("/", h.createComment)
			r.Get("/", h.getPostComments)
			r.Put("/{comment_id}", h.updateComment)
		})
	})
	return r
}

func NewLogger() *zap.Logger {
	rawJSON := []byte(`{
	  "level": "error",
	  "encoding": "json",
	  "outputPaths": ["stdout"],
	  "errorOutputPaths": ["stderr"],
	  "encoderConfig": {
	    "messageKey": "message",
	    "levelKey": "level",
	    "levelEncoder": "lowercase"
	  }
	}`)
	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	logger := zap.Must(cfg.Build())
	defer logger.Sync()
	return logger
}
