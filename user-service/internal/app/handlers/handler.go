package handlers

import (
	"encoding/json"
	"github.com/HeadGardener/blog-app/user-service/internal/app/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
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
		errLogger: newLogger(),
	}
}

func InitRoutes(h *Handler) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Route("/api/users", func(r chi.Router) {
		r.Post("/sign-up", h.signUp)
		r.Get("/", h.signIn)
	})

	return r
}

func newLogger() *zap.Logger {
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
