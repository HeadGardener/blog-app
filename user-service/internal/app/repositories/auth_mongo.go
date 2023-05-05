package repositories

import (
	"context"
	"github.com/HeadGardener/blog-app/user-service/internal/app/models"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type AuthRepository struct {
	db *mongo.Collection
}

func NewAuthRepository(db *mongo.Collection) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) Create(ctx context.Context, user models.User) (string, error) {
	insertCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := r.db.InsertOne(insertCtx, user)
	if err != nil {
		return "", err
	}

	return user.ID, nil
}
