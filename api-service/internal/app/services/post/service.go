package postService

import (
	"context"
	"github.com/HeadGardener/blog-app/api-service/internal/app/models"
	"github.com/HeadGardener/blog-app/api-service/pkg/client"
	"net/http"
	"time"
)

type service struct {
	base     client.Client
	Resource string
}

func NewPostService(baseURL, resource string) PostService {
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

type PostService interface {
	CreatePost(ctx context.Context, postInput models.CreatePostInput) (string, error)
	GetPostByID(ctx context.Context, id string) (models.Post, error)
	GetPosts(ctx context.Context, userID string, postsAmount string) ([]models.Post, error)
	UpdatePost(ctx context.Context, postInput models.UpdatePostInput) (models.Post, error)
	DeletePost(ctx context.Context, postID, userID string) (models.Post, error)
}
