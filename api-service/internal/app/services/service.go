package services

import (
	"github.com/HeadGardener/blog-app/api-service/configs"
	commentService "github.com/HeadGardener/blog-app/api-service/internal/app/services/comment"
	postService "github.com/HeadGardener/blog-app/api-service/internal/app/services/post"
	userService "github.com/HeadGardener/blog-app/api-service/internal/app/services/user"
)

type Service struct {
	userService.UserService
	postService.PostService
	commentService.CommentService
}

func NewService(config configs.ServiceConfig) *Service {
	return &Service{
		UserService:    userService.NewUserService(config.UserServiceURL, "users"),
		PostService:    postService.NewPostService(config.PostServiceURL, "post"),
		CommentService: commentService.NewCommentService(config.CommentServiceURL, "comment"),
	}
}
