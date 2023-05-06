package services

import (
	"context"
	"github.com/HeadGardener/blog-app/post-service/internal/app/models"
	"github.com/HeadGardener/blog-app/post-service/internal/app/repositories"
)

type PostInterface interface {
	Create(ctx context.Context, postInput models.CreatePostInput) (string, error)
}

type Service struct {
	PostInterface
}

func NewService(repository *repositories.Repository) *Service {
	return &Service{
		PostInterface: NewPostService(repository),
	}
}
