package configs

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

type ServiceConfig struct {
	UserServiceURL string
}

func NewServiceConfig(path string) (*ServiceConfig, error) {
	err := godotenv.Load(path)
	if err != nil {
		return nil, err
	}

	userServiceURL := os.Getenv("USER_SERVICE_URL")
	if userServiceURL == "" {
		return nil, errors.New("server port is empty")
	}

	return &ServiceConfig{
		UserServiceURL: userServiceURL,
	}, nil
}
