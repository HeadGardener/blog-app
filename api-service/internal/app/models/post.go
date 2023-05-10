package models

import (
	"errors"
	"time"
)

type Post struct {
	ID     string    `json:"id"`
	UserID string    `json:"userID"`
	Title  string    `json:"title"`
	Body   string    `json:"body"`
	Date   time.Time `json:"date"`
}

type CreatePostInput struct {
	UserID string `json:"userID"`
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
	if p.Body == "" {
		return errors.New("post body can't be empty")
	}

	return nil
}
