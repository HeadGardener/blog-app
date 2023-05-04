package user_service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/HeadGardener/api-service/internal/app/models"
	"github.com/HeadGardener/api-service/pkg/client"
	"net/http"
	"time"
)

type service struct {
	base     client.Client
	Resource string
}

func NewUserService(baseURL, resource string) UserService {
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

type UserService interface {
	Create(ctx context.Context, userInput models.CreateUserInput) (string, error)
}

type errResponse struct {
	Message string `json:"message"`
}

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
		var errMsg errResponse
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
		return "", fmt.Errorf("emfiefiefme")
	}

	return uuid.ID, nil
}
