package services

import (
	"context"
	"github.com/HeadGardener/blog-app/user-service/internal/app/models"
	"github.com/HeadGardener/blog-app/user-service/internal/app/repositories"
)

type Authorization interface {
	Create(ctx context.Context, userInput models.CreateUserInput) (string, error)
}

type UserInterface interface {
	GetUser(ctx context.Context, userInput models.LogUserInput) (models.User, error)
}

type Service struct {
	Authorization
	UserInterface
}

func NewService(repository *repositories.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repository),
		UserInterface: NewUserService(repository),
	}
}
