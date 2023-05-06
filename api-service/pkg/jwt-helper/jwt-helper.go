package jwt_helper

import (
	"errors"
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

func ParseToken(accessToken string) (models.UserAttributes, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return models.UserAttributes{}, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return models.UserAttributes{}, errors.New("token claims are not of type *tokenClaims")
	}

	return models.UserAttributes{
		ID:    claims.UserID,
		Email: claims.Email,
	}, nil
}
