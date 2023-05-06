package repositories

import (
	"context"
	"github.com/HeadGardener/blog-app/user-service/internal/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type UserRepository struct {
	db *mongo.Collection
}

func NewUserRepository(db *mongo.Collection) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUser(ctx context.Context, user models.User) (models.User, error) {
	insertCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	result := r.db.FindOne(insertCtx, bson.D{
		{"email", user.Email},
		{"passwordhash", user.PasswordHash}})

	if err := result.Decode(&user); err != nil {
		return models.User{}, err
	}

	return user, nil
}
