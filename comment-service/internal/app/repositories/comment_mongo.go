package repositories

import (
	"context"
	"github.com/HeadGardener/blog-app/comment-service/internal/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type CommentRepository struct {
	db *mongo.Collection
}

func NewCommentRepository(db *mongo.Collection) *CommentRepository {
	return &CommentRepository{db: db}
}

func (r *CommentRepository) CreateComment(ctx context.Context, comment models.Comment) (string, error) {
	insertCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := r.db.InsertOne(insertCtx, comment)
	if err != nil {
		return "", err
	}

	return comment.ID, nil
}

func (r *CommentRepository) GetByID(ctx context.Context, commentID string) (models.Comment, error) {
	var comment models.Comment

	insertCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	result := r.db.FindOne(insertCtx, bson.D{
		{"id", commentID}})

	if err := result.Decode(&comment); err != nil {
		return models.Comment{}, err
	}

	return comment, nil
}

func (r *CommentRepository) GetComments(ctx context.Context, postID string, commentsAmount int) ([]models.Comment, error) {
	var comments []models.Comment

	insertCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	opts := options.Find().SetLimit(int64(commentsAmount)).SetSort(bson.D{{"date", 1}})
	cur, err := r.db.Find(insertCtx, bson.D{{
		"post_id", postID,
	}}, opts)
	if err != nil {
		return nil, err
	}

	if err := cur.All(insertCtx, &comments); err != nil {
		return nil, err
	}

	return comments, nil
}

func (r *CommentRepository) UpdateComment(ctx context.Context, comment models.Comment) (models.Comment, error) {
	var updatedComment models.Comment

	insertCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	opts := options.FindOneAndUpdate().SetUpsert(false)
	filters := bson.D{{"id", comment.ID}, {"user_id", comment.UserID}}
	update := bson.D{{"$set", bson.D{{"body", comment.Body}}}}
	if err := r.db.FindOneAndUpdate(insertCtx, filters, update, opts).Decode(&updatedComment); err != nil {
		return models.Comment{}, err
	}

	return updatedComment, nil
}
