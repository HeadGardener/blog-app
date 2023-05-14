package repositories

import (
	"context"
	"github.com/HeadGardener/blog-app/post-service/internal/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (r *PostRepository) GetPosts(ctx context.Context, userID string, postsAmount int) ([]models.Post, error) {
	var posts []models.Post

	insertCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	opts := options.Find().SetLimit(int64(postsAmount)).SetSort(bson.D{{"date", 1}})
	cur, err := r.db.Find(insertCtx, bson.D{{
		"user_id", userID,
	}}, opts)
	if err != nil {
		return nil, err
	}

	if err := cur.All(insertCtx, &posts); err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *PostRepository) UpdatePost(ctx context.Context, post models.Post) (models.Post, error) {
	var updatedPost models.Post

	insertCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	opts := options.FindOneAndUpdate().SetUpsert(false)
	filters := bson.D{{"id", post.ID}, {"user_id", post.UserID}}
	update := bson.D{{"$set", bson.D{{"title", post.Title}, {"body", post.Body}}}}
	if err := r.db.FindOneAndUpdate(insertCtx, filters, update, opts).Decode(&updatedPost); err != nil {
		return models.Post{}, err
	}

	return updatedPost, nil
}

func (r *PostRepository) DeletePost(ctx context.Context, postID, userID string) (models.Post, error) {
	var deletedPost models.Post

	insertCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filters := bson.D{{"id", postID}, {"userid", userID}}
	if err := r.db.FindOneAndDelete(insertCtx, filters).Decode(&deletedPost); err != nil {
		return models.Post{}, err
	}

	return deletedPost, nil
}
