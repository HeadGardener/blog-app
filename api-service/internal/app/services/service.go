package services

import (
	"github.com/HeadGardener/blog-app/api-service/configs"
	postService "github.com/HeadGardener/blog-app/api-service/internal/app/services/post"
	userService "github.com/HeadGardener/blog-app/api-service/internal/app/services/user"
)

type Service struct {
	userService.UserService
	postService.PostService
}

func NewService(config configs.ServiceConfig) *Service {
	return &Service{
		UserService: userService.NewUserService(config.UserServiceURL, "users"),
		PostService: postService.NewPostService(config.PostServiceURL, "post"),
	}
}
