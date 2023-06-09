package commentService

import (
	"context"
	"github.com/HeadGardener/blog-app/api-service/internal/app/models"
	"github.com/HeadGardener/blog-app/api-service/pkg/client"
	"net/http"
	"time"
)

type CommentService interface {
	CreateComment(ctx context.Context, commentInput models.CreateCommentInput) (string, error)
	GetComments(ctx context.Context, postID, commentsAmount string) ([]models.Comment, error)
	UpdateComment(ctx context.Context, commentInput models.UpdateCommentInput) (models.Comment, error)
	DeleteAllPostComments(ctx context.Context, postID string) error
	DeleteComment(ctx context.Context, commentID, userID string) (models.Comment, error)
}

type service struct {
	base     client.Client
	Resource string
}

func NewCommentService(baseURL, resource string) CommentService {
	service := service{
		Resource: resource,
		base: client.Client{
			BaseURL: baseURL,
			HTTPClient: &http.Client{
				Timeout: 10 * time.Second,
			},
		},
	}
	return &service
}
