package services

import (
	"context"
	"github.com/HeadGardener/user-service/internal/app/models"
	"github.com/HeadGardener/user-service/internal/app/repositories"
)

type Authorization interface {
	Create(ctx context.Context, userInput models.CreateUserInput) (string, error)
}

type Service struct {
	Authorization
}

func NewService(repository *repositories.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repository),
	}
}
