package configs

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

type ServerConfig struct {
	ServerPort string
}

func NewServerConfig(path string) (*ServerConfig, error) {
	err := godotenv.Load(path)
	if err != nil {
		return nil, err
	}

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		return nil, errors.New("server port is empty")
	}

	return &ServerConfig{
		ServerPort: serverPort,
	}, nil
}
