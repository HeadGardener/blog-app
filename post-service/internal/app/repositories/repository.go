package repositories

import (
	"context"
	"github.com/HeadGardener/blog-app/post-service/internal/app/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostInterface interface {
	Create(ctx context.Context, post models.Post) (string, error)
}

type Repository struct {
	PostInterface
}

func NewRepository(db *mongo.Collection) *Repository {
	return &Repository{
		PostInterface: NewPostRepository(db),
	}
}
