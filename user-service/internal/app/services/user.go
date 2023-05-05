package services

import (
	"context"
	"github.com/HeadGardener/blog-app/user-service/internal/app/models"
	"github.com/HeadGardener/blog-app/user-service/internal/app/repositories"
)

type UserService struct {
	repository *repositories.Repository
}

func NewUserService(repository *repositories.Repository) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) GetUser(ctx context.Context, userInput models.LogUserInput) (models.User, error) {
	user := models.User{
		Email:        userInput.Email,
		PasswordHash: getPasswordHash(userInput.Password),
	}

	return s.repository.UserInterface.GetUser(ctx, user)
}
