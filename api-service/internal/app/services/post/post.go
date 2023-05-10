package postService

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/HeadGardener/blog-app/api-service/internal/app/models"
	"github.com/HeadGardener/blog-app/api-service/pkg/responses"
	"net/http"
	"time"
)

func (s *service) CreatePost(ctx context.Context, postInput models.CreatePostInput) (string, error) {
	url := fmt.Sprintf("%s/%s/", s.base.BaseURL, s.Resource)

	dataBytes, err := json.Marshal(postInput)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(dataBytes))
	if err != nil {
		return "", fmt.Errorf("failed to create request: error: %w", err)
	}

	reqCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	req = req.WithContext(reqCtx)
	response, err := s.base.SendRequest(req)
	defer response.Body.Close()
	if err != nil {
		return "", fmt.Errorf("failed to send request: error: %w", err)
	}

	if response.StatusCode != http.StatusCreated {
		var errMsg responses.ErrResponse
		if err := json.NewDecoder(response.Body).Decode(&errMsg); err != nil {
			return "", fmt.Errorf("unpredictable error")
		}

		return "", fmt.Errorf("%s", errMsg.Message)
	}

	type uuidResponse struct {
		ID string `json:"id"`
	}

	var uuid uuidResponse
	if err := json.NewDecoder(response.Body).Decode(&uuid); err != nil {
		return "", fmt.Errorf("error while decoding response: error: %w", err)
	}

	return uuid.ID, nil
}

func (s *service) GetPostByID(ctx context.Context, postID string) (models.Post, error) {
	url := fmt.Sprintf("%s/%s/%s", s.base.BaseURL, s.Resource, postID)

	req, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer([]byte("")))
	if err != nil {
		return models.Post{}, fmt.Errorf("failed to create request: error: %w", err)
	}

	reqCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	req = req.WithContext(reqCtx)
	response, err := s.base.SendRequest(req)
	defer response.Body.Close()
	if err != nil {
		return models.Post{}, fmt.Errorf("failed to send request: error: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		var errMsg responses.ErrResponse
		if err := json.NewDecoder(response.Body).Decode(&errMsg); err != nil {
			return models.Post{}, fmt.Errorf("unpredictable error")
		}

		return models.Post{}, fmt.Errorf("%s", errMsg.Message)
	}

	var post models.Post
	if err := json.NewDecoder(response.Body).Decode(&post); err != nil {
		return models.Post{}, fmt.Errorf("error while decoding response: error: %w", err)
	}

	return post, nil
}

func (s *service) GetPosts(ctx context.Context, userID string, postsAmount string) ([]models.Post, error) {
	url := fmt.Sprintf("%s/%s/?user_id=%s&amount=%s", s.base.BaseURL, s.Resource, userID, postsAmount)

	req, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer([]byte("")))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: error: %w", err)
	}

	reqCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	req = req.WithContext(reqCtx)
	response, err := s.base.SendRequest(req)
	defer response.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to send request: error: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		var errMsg responses.ErrResponse
		if err := json.NewDecoder(response.Body).Decode(&errMsg); err != nil {
			return nil, fmt.Errorf("unpredictable error")
		}

		return nil, fmt.Errorf("%s", errMsg.Message)
	}

	var posts []models.Post
	if err := json.NewDecoder(response.Body).Decode(&posts); err != nil {
		return nil, fmt.Errorf("error while decoding response: error: %w", err)
	}

	return posts, nil
}

func (s *service) UpdatePost(ctx context.Context, postInput models.UpdatePostInput) (models.Post, error) {
	url := fmt.Sprintf("%s/%s/", s.base.BaseURL, s.Resource)

	dataBytes, err := json.Marshal(postInput)
	if err != nil {
		return models.Post{}, err
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(dataBytes))
	if err != nil {
		return models.Post{}, fmt.Errorf("failed to create request: error: %w", err)
	}

	reqCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	req = req.WithContext(reqCtx)
	response, err := s.base.SendRequest(req)
	defer response.Body.Close()
	if err != nil {
		return models.Post{}, fmt.Errorf("failed to send request: error: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		var errMsg responses.ErrResponse
		if err := json.NewDecoder(response.Body).Decode(&errMsg); err != nil {
			return models.Post{}, fmt.Errorf("unpredictable error")
		}

		return models.Post{}, fmt.Errorf("%s", errMsg.Message)
	}

	var post models.Post
	if err := json.NewDecoder(response.Body).Decode(&post); err != nil {
		return models.Post{}, fmt.Errorf("error while decoding response: error: %w", err)
	}
	return post, nil
}
