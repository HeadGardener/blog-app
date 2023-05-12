package services

import (
	"context"
	"github.com/HeadGardener/blog-app/comment-service/internal/app/models"
	"github.com/HeadGardener/blog-app/comment-service/internal/app/repositories"
)

type CommentInterface interface {
	Create(ctx context.Context, commentInput models.CreateCommentInput) (string, error)
}

type Service struct {
	CommentInterface
}

func NewService(repository *repositories.Repository) *Service {
	return &Service{
		CommentInterface: NewCommentService(repository),
	}
}
