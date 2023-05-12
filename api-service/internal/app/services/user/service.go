package userService

import (
	"context"
	"github.com/HeadGardener/blog-app/api-service/internal/app/models"
	"github.com/HeadGardener/blog-app/api-service/pkg/client"
	"net/http"
	"time"
)

type UserService interface {
	Create(ctx context.Context, userInput models.CreateUserInput) (string, error)
	GetUser(ctx context.Context, userInput models.LogUserInput) (models.User, error)
}

type service struct {
	base     client.Client
	Resource string
}

func NewUserService(baseURL, resource string) UserService {
	service := service{
		Resource: resource,
		base: client.Client{
			BaseURL: baseURL,
			HTTPClient: &http.Client{
				Timeout: 10 * time.Second,
			},
		},
	}
	return &service
}
