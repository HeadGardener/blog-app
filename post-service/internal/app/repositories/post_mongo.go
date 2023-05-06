package repositories

import (
	"context"
	"github.com/HeadGardener/blog-app/post-service/internal/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type PostRepository struct {
	db *mongo.Collection
}

func NewPostRepository(db *mongo.Collection) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) Create(ctx context.Context, post models.Post) (string, error) {
	insertCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := r.db.InsertOne(insertCtx, post)
	if err != nil {
		return "", err
	}

	return post.ID, nil
}

func (r *PostRepository) GetByID(ctx context.Context, postID string) (models.Post, error) {
	var post models.Post

	insertCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	result := r.db.FindOne(insertCtx, bson.D{
		{"id", postID}})

	if err := result.Decode(&post); err != nil {
		return models.Post{}, err
	}

	return post, nil
}
