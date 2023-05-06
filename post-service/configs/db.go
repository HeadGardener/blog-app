package configs

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

type DBConfig struct {
	Host   string
	Port   string
	DBName string
}

func NewDBConfig(path string) (*DBConfig, error) {
	err := godotenv.Load(path)
	if err != nil {
		return nil, err
	}

	host := os.Getenv("DB_HOST")
	if host == "" {
		return nil, errors.New("db host is empty")
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		return nil, errors.New("db name is empty")
	}

	name := os.Getenv("DB_NAME")
	if name == "" {
		return nil, errors.New("db name is empty")
	}

	return &DBConfig{
		Host:   host,
		Port:   port,
		DBName: name,
	}, nil
}
