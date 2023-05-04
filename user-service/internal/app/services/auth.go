package services

import (
	"context"
	"crypto/sha1"
	"fmt"
	"github.com/HeadGardener/user-service/internal/app/models"
	"github.com/HeadGardener/user-service/internal/app/repositories"
	"github.com/google/uuid"
	"time"
)

const (
	salt      = "qetuoadgjlzcbmwryipsfhkxvn"
	tokenTTL  = 15 * time.Minute
	secretKey = "qazwsxedcrfvtgbyhnujm"
)

type AuthService struct {
	repository *repositories.Repository
}

func NewAuthService(repository *repositories.Repository) *AuthService {
	return &AuthService{repository: repository}
}

func (s *AuthService) Create(ctx context.Context, userInput models.CreateUserInput) (string, error) {
	user := models.User{
		ID:           uuid.NewString(),
		Name:         userInput.Name,
		Surname:      userInput.Surname,
		Email:        userInput.Email,
		PasswordHash: getPasswordHash(userInput.Password),
	}

	return s.repository.Create(ctx, user)
}

func getPasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
