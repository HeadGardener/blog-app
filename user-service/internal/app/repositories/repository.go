package repositories

import (
	"context"
	"github.com/HeadGardener/blog-app/user-service/internal/app/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type Authorization interface {
	Create(ctx context.Context, user models.User) (string, error)
}

type UserInterface interface {
	GetUser(ctx context.Context, user models.User) (models.User, error)
}

type Repository struct {
	Authorization
	UserInterface
}

func NewRepository(db *mongo.Collection) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
		UserInterface: NewUserRepository(db),
	}
}
