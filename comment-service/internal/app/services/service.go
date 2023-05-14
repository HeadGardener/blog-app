package services

import (
	"context"
	"github.com/HeadGardener/blog-app/comment-service/internal/app/models"
	"github.com/HeadGardener/blog-app/comment-service/internal/app/repositories"
)

type CommentInterface interface {
	CreateComment(ctx context.Context, commentInput models.CreateCommentInput) (string, error)
	GetComments(ctx context.Context, postID string, amount int) ([]models.Comment, error)
	UpdateComment(ctx context.Context, commentInput models.UpdateCommentInput) (models.Comment, error)
}

type Service struct {
	CommentInterface
}

func NewService(repository *repositories.Repository) *Service {
	return &Service{
		CommentInterface: NewCommentService(repository),
	}
}
