package models

import (
	"errors"
	"time"
)

type Post struct {
	ID     string    `json:"id" bson:"id"`
	UserID string    `json:"user_id" bson:"user_id"`
	Title  string    `json:"title" bson:"title"`
	Body   string    `json:"body" bson:"body"`
	Date   time.Time `json:"date" bson:"date"`
}

type CreatePostInput struct {
	UserID string `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type UpdatePostInput struct {
	ID     string  `json:"id"`
	UserID string  `json:"user_id"`
	Title  *string `json:"title"`
	Body   *string `json:"body"`
}

func (p *CreatePostInput) Validate() error {
	if p.UserID == "" {
		return errors.New("invalid user id")
	}
	if p.Body == "" {
		return errors.New("post body can't be empty")
	}

	return nil
}

func (p *UpdatePostInput) Validate() error {
	if p.ID == "" || p.UserID == "" {
		return errors.New("invalid or empty ids values")
	}
	return nil
}

func (p *UpdatePostInput) ToPost(post *Post) {
	if p.Title != nil && post.Title != *p.Title {
		post.Title = *p.Title
	}

	if p.Body != nil && *p.Body != "" && post.Body != *p.Body {
		post.Body = *p.Body
	}
}
