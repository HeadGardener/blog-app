package repositories

import (
	"context"
	"github.com/HeadGardener/blog-app/post-service/internal/app/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostInterface interface {
	Create(ctx context.Context, post models.Post) (string, error)
	GetByID(ctx context.Context, postID string) (models.Post, error)
	GetPosts(ctx context.Context, userID string, postsAmount int) ([]models.Post, error)
	UpdatePost(ctx context.Context, postInput models.Post) (models.Post, error)
	DeletePost(ctx context.Context, postID, userID string) (models.Post, error)
}

type Repository struct {
	PostInterface
}

func NewRepository(db *mongo.Collection) *Repository {
	return &Repository{
		PostInterface: NewPostRepository(db),
	}
}
