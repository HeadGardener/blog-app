package services

import (
	"context"
	"errors"
	"github.com/HeadGardener/blog-app/post-service/internal/app/models"
	"github.com/HeadGardener/blog-app/post-service/internal/app/repositories"
	"github.com/google/uuid"
	"time"
)

type PostService struct {
	repository *repositories.Repository
}

func NewPostService(repository *repositories.Repository) *PostService {
	return &PostService{repository: repository}
}

func (s *PostService) CreatePost(ctx context.Context, postInput models.CreatePostInput) (string, error) {
	post := models.Post{
		ID:     uuid.NewString(),
		UserID: postInput.UserID,
		Title:  postInput.Title,
		Body:   postInput.Body,
		Date:   time.Now(),
	}

	return s.repository.PostInterface.Create(ctx, post)
}

func (s *PostService) GetByID(ctx context.Context, postID string) (models.Post, error) {
	return s.repository.PostInterface.GetByID(ctx, postID)
}

func (s *PostService) GetPosts(ctx context.Context, userID string, postsAmount int) ([]models.Post, error) {
	return s.repository.PostInterface.GetPosts(ctx, userID, postsAmount)
}

func (s *PostService) UpdatePost(ctx context.Context, postInput models.UpdatePostInput) (models.Post, error) {
	post, err := s.repository.GetByID(ctx, postInput.ID)
	if err != nil {
		return models.Post{}, err
	}
	
	if post.UserID != postInput.UserID {
		return models.Post{}, errors.New("this is not your post")
	}

	postInput.ToPost(&post)

	return s.repository.PostInterface.UpdatePost(ctx, post)
}

func (s *PostService) DeletePost(ctx context.Context, postID, userID string) (models.Post, error) {
	return s.repository.PostInterface.DeletePost(ctx, postID, userID)
}
