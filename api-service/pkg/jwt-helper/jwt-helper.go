package jwt_helper

import (
	"github.com/HeadGardener/blog-app/api-service/internal/app/models"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	tokenTTL  = 15 * time.Minute
	secretKey = "qazwsxedcrfvtgbyhnujm"
)

type tokenClaims struct {
	jwt.RegisteredClaims
	UserID string `json:"user_id"`
	Email  string `json:"email"`
}

func GenerateToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		user.ID,
		user.Email,
	})

	return token.SignedString([]byte(secretKey))
}
