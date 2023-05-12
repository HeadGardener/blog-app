package services

import (
	"context"
	"github.com/HeadGardener/blog-app/comment-service/internal/app/models"
	"github.com/HeadGardener/blog-app/comment-service/internal/app/repositories"
	"github.com/google/uuid"
	"time"
)

type CommentService struct {
	repository *repositories.Repository
}

func NewCommentService(repository *repositories.Repository) *CommentService {
	return &CommentService{repository: repository}
}

func (s *CommentService) Create(ctx context.Context, commentInput models.CreateCommentInput) (string, error) {
	comment := models.Comment{
		ID:     uuid.NewString(),
		PostID: commentInput.PostID,
		UserID: commentInput.UserID,
		Body:   commentInput.Body,
		Date:   time.Now(),
	}

	return s.repository.CommentInterface.Create(ctx, comment)
}
