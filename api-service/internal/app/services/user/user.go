package userService

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

func (s *service) Create(ctx context.Context, userInput models.CreateUserInput) (string, error) {
	url := fmt.Sprintf("%s/%s/sign-up", s.base.BaseURL, s.Resource)

	dataBytes, err := json.Marshal(userInput)
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

func (s *service) GetUser(ctx context.Context, userInput models.LogUserInput) (models.User, error) {
	url := fmt.Sprintf("%s/%s/", s.base.BaseURL, s.Resource)

	dataBytes, err := json.Marshal(userInput)
	if err != nil {
		return models.User{}, err
	}

	req, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(dataBytes))
	if err != nil {
		return models.User{}, fmt.Errorf("failed to create request: error: %w", err)
	}

	reqCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	req = req.WithContext(reqCtx)
	response, err := s.base.SendRequest(req)
	defer response.Body.Close()
	if err != nil {
		return models.User{}, fmt.Errorf("failed to send request: error: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		var errMsg responses.ErrResponse
		if err := json.NewDecoder(response.Body).Decode(&errMsg); err != nil {
			return models.User{}, fmt.Errorf("unpredictable error")
		}

		return models.User{}, fmt.Errorf("%s", errMsg.Message)
	}

	var user models.User

	if err := json.NewDecoder(response.Body).Decode(&user); err != nil {
		return models.User{}, fmt.Errorf("error while decoding response: error: %w", err)
	}

	return user, nil
}
