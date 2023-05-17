package services

import (
	"context"
	"errors"
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

func (s *CommentService) CreateComment(ctx context.Context, commentInput models.CreateCommentInput) (string, error) {
	comment := models.Comment{
		ID:     uuid.NewString(),
		PostID: commentInput.PostID,
		UserID: commentInput.UserID,
		Body:   commentInput.Body,
		Date:   time.Now(),
	}

	return s.repository.CommentInterface.CreateComment(ctx, comment)
}

func (s *CommentService) GetComments(ctx context.Context, postID string, amount int) ([]models.Comment, error) {
	return s.repository.CommentInterface.GetComments(ctx, postID, amount)
}

func (s *CommentService) UpdateComment(ctx context.Context, commentInput models.UpdateCommentInput) (models.Comment, error) {
	comment, err := s.repository.CommentInterface.GetByID(ctx, commentInput.ID)
	if err != nil {
		return models.Comment{}, err
	}

	if comment.UserID != commentInput.UserID {
		return models.Comment{}, errors.New("this is not your comment")
	}

	commentInput.ToComment(&comment)

	return s.repository.CommentInterface.UpdateComment(ctx, comment)
}

func (s *CommentService) DeleteAllPostComments(ctx context.Context, postID string) error {
	return s.repository.CommentInterface.DeleteAllPostComments(ctx, postID)
}

func (s *CommentService) DeleteComment(ctx context.Context, commentID, userID string) (models.Comment, error) {
	return s.repository.CommentInterface.DeleteComment(ctx, commentID, userID)
}
