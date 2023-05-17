package commentService

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

func (s *service) CreateComment(ctx context.Context, commentInput models.CreateCommentInput) (string, error) {
	url := fmt.Sprintf("%s/%s/", s.base.BaseURL, s.Resource)

	dataBytes, err := json.Marshal(commentInput)
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

func (s *service) GetComments(ctx context.Context, postID, commentsAmount string) ([]models.Comment, error) {
	url := fmt.Sprintf("%s/%s/?post_id=%s&amount=%s", s.base.BaseURL, s.Resource, postID, commentsAmount)

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

	var comments []models.Comment
	if err := json.NewDecoder(response.Body).Decode(&comments); err != nil {
		return nil, fmt.Errorf("error while decoding response: error: %w", err)
	}

	return comments, nil
}

func (s *service) UpdateComment(ctx context.Context, commentInput models.UpdateCommentInput) (models.Comment, error) {
	url := fmt.Sprintf("%s/%s/", s.base.BaseURL, s.Resource)

	dataBytes, err := json.Marshal(commentInput)
	if err != nil {
		return models.Comment{}, err
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(dataBytes))
	if err != nil {
		return models.Comment{}, fmt.Errorf("failed to create request: error: %w", err)
	}

	reqCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	req = req.WithContext(reqCtx)
	response, err := s.base.SendRequest(req)
	defer response.Body.Close()
	if err != nil {
		return models.Comment{}, fmt.Errorf("failed to send request: error: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		var errMsg responses.ErrResponse
		if err := json.NewDecoder(response.Body).Decode(&errMsg); err != nil {
			return models.Comment{}, fmt.Errorf("unpredictable error")
		}

		return models.Comment{}, fmt.Errorf("%s", errMsg.Message)
	}

	var comment models.Comment
	if err := json.NewDecoder(response.Body).Decode(&comment); err != nil {
		return models.Comment{}, fmt.Errorf("error while decoding response: error: %w", err)
	}
	return comment, nil
}

func (s *service) DeleteAllPostComments(ctx context.Context, postID string) error {
	url := fmt.Sprintf("%s/%s/post/%s", s.base.BaseURL, s.Resource, postID)

	req, err := http.NewRequest(http.MethodDelete, url, bytes.NewBuffer([]byte("")))
	if err != nil {
		return fmt.Errorf("failed to create request: error: %w", err)
	}

	reqCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	req = req.WithContext(reqCtx)
	response, err := s.base.SendRequest(req)
	defer response.Body.Close()
	if err != nil {
		return fmt.Errorf("failed to send request: error: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		var errMsg responses.ErrResponse
		if err := json.NewDecoder(response.Body).Decode(&errMsg); err != nil {
			return fmt.Errorf("unpredictable error")
		}

		return fmt.Errorf("%s", errMsg.Message)
	}

	return nil
}

func (s *service) DeleteComment(ctx context.Context, commentID, userID string) (models.Comment, error) {
	url := fmt.Sprintf("%s/%s/%s?user_id=%s", s.base.BaseURL, s.Resource, commentID, userID)

	req, err := http.NewRequest(http.MethodDelete, url, bytes.NewBuffer([]byte("")))
	if err != nil {
		return models.Comment{}, fmt.Errorf("failed to create request: error: %w", err)
	}

	reqCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	req = req.WithContext(reqCtx)
	response, err := s.base.SendRequest(req)
	defer response.Body.Close()
	if err != nil {
		return models.Comment{}, fmt.Errorf("failed to send request: error: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		var errMsg responses.ErrResponse
		if err := json.NewDecoder(response.Body).Decode(&errMsg); err != nil {
			return models.Comment{}, fmt.Errorf("unpredictable error")
		}

		return models.Comment{}, fmt.Errorf("%s", errMsg.Message)
	}

	var comment models.Comment
	if err := json.NewDecoder(response.Body).Decode(&comment); err != nil {
		return models.Comment{}, fmt.Errorf("error while decoding response: error: %w", err)
	}
	return comment, nil
}
