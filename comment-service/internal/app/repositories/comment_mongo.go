package repositories

import (
	"context"
	"github.com/HeadGardener/blog-app/comment-service/internal/app/models"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type CommentRepository struct {
	db *mongo.Collection
}

func NewCommentRepository(db *mongo.Collection) *CommentRepository {
	return &CommentRepository{db: db}
}

func (r *CommentRepository) Create(ctx context.Context, comment models.Comment) (string, error) {
	insertCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := r.db.InsertOne(insertCtx, comment)
	if err != nil {
		return "", err
	}

	return comment.ID, nil
}
