package repositories

import (
	"context"
	"github.com/HeadGardener/blog-app/comment-service/internal/app/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type CommentInterface interface {
	CreateComment(ctx context.Context, comment models.Comment) (string, error)
	GetByID(ctx context.Context, commentID string) (models.Comment, error)
	GetComments(ctx context.Context, postID string, amount int) ([]models.Comment, error)
	UpdateComment(ctx context.Context, comment models.Comment) (models.Comment, error)
}

type Repository struct {
	CommentInterface
}

func NewRepository(db *mongo.Collection) *Repository {
	return &Repository{
		CommentInterface: NewCommentRepository(db),
	}
}
