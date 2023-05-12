package repositories

import (
	"context"
	"github.com/HeadGardener/blog-app/comment-service/internal/app/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type CommentInterface interface {
	Create(ctx context.Context, comment models.Comment) (string, error)
}

type Repository struct {
	CommentInterface
}

func NewRepository(db *mongo.Collection) *Repository {
	return &Repository{
		CommentInterface: NewCommentRepository(db),
	}
}
