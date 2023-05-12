package models

import (
	"errors"
	"strings"
)

type Comment struct {
	ID     string `json:"id"`
	PostID string `json:"post_id"`
	UserID string `json:"user_id"`
	Body   string `json:"body"`
	Date   string `json:"date"`
}

type CreateCommentInput struct {
	PostID string `json:"post_id"`
	UserID string `json:"user_id"`
	Body   string `json:"body"`
}

func (c *CreateCommentInput) Validate() error {
	if c.PostID == "" {
		return errors.New("invalid or empty post id")
	}

	if strings.ReplaceAll(c.Body, " ", "") == "" {
		return errors.New("comment body can't be empty")
	}

	return nil
}
