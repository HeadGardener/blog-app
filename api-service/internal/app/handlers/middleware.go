package handlers

import (
	"context"
	"errors"
	"github.com/HeadGardener/blog-app/api-service/internal/app/models"
	jwt_helper "github.com/HeadGardener/blog-app/api-service/pkg/jwt-helper"
	"github.com/HeadGardener/blog-app/api-service/pkg/responses"
	"net/http"
	"strings"
)

const userCtx = "userAtr"

var mwLogger = NewLogger()

func IdentifyUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")

		if header == "" {
			responses.NewErrResponse(w, http.StatusUnauthorized, "empty auth header", mwLogger)
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			responses.NewErrResponse(w, http.StatusUnauthorized, "invalid auth header", mwLogger)
			return
		}

		if headerParts[0] != "Bearer" {
			responses.NewErrResponse(w, http.StatusUnauthorized, "invalid auth header", mwLogger)
			return
		}

		if len(headerParts[1]) == 0 {
			responses.NewErrResponse(w, http.StatusUnauthorized, "jwt token is empty", mwLogger)
			return
		}

		userAttributes, err := jwt_helper.ParseToken(headerParts[1])
		if err != nil {
			responses.NewErrResponse(w, http.StatusUnauthorized, err.Error(), mwLogger)
			return
		}

		ctx := context.WithValue(r.Context(), userCtx, userAttributes)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUserID(r *http.Request) (string, error) {
	workerCtxValue := r.Context().Value(userCtx)
	workerAttributes, ok := workerCtxValue.(models.UserAttributes)
	if !ok {
		return "", errors.New("workerCtx value is not of type WorkerAttributes")
	}

	return workerAttributes.ID, nil
}
