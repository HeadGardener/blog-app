package post_service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/HeadGardener/blog-app/api-service/internal/app/models"
	"github.com/HeadGardener/blog-app/api-service/pkg/client"
	"github.com/HeadGardener/blog-app/api-service/pkg/responses"
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
}

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
