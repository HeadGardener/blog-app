package configs

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

type ServiceConfig struct {
	UserServiceURL    string
	PostServiceURL    string
	CommentServiceURL string
}

func NewServiceConfig(path string) (*ServiceConfig, error) {
	err := godotenv.Load(path)
	if err != nil {
		return nil, err
	}

	userServiceURL := os.Getenv("USER_SERVICE_URL")
	if userServiceURL == "" {
		return nil, errors.New("USER_SERVICE_URL is empty")
	}

	postServiceURL := os.Getenv("POST_SERVICE_URL")
	if userServiceURL == "" {
		return nil, errors.New("POST_SERVICE_URL is empty")
	}

	commentServiceURL := os.Getenv("COMMENT_SERVICE_URL")
	if userServiceURL == "" {
		return nil, errors.New("COMMENT_SERVICE_URL is empty")
	}

	return &ServiceConfig{
		UserServiceURL:    userServiceURL,
		PostServiceURL:    postServiceURL,
		CommentServiceURL: commentServiceURL,
	}, nil
}
