package repositories

import (
	"context"
	"github.com/HeadGardener/user-service/internal/app/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type Authorization interface {
	Create(ctx context.Context, user models.User) (string, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *mongo.Collection) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
	}
}
