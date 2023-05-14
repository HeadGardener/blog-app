package models

import (
	"errors"
	"strings"
	"time"
)

type Comment struct {
	ID     string    `json:"id" bson:"id"`
	PostID string    `json:"post_id" bson:"post_id"`
	UserID string    `json:"user_id" bson:"user_id"`
	Body   string    `json:"body" bson:"body"`
	Date   time.Time `json:"date" bson:"date"`
}

type CreateCommentInput struct {
	PostID string `json:"post_id"`
	UserID string `json:"user_id"`
	Body   string `json:"body"`
}

type UpdateCommentInput struct {
	ID     string  `json:"id"`
	UserID string  `json:"user_id"`
	Body   *string `json:"body"`
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

func (c *UpdateCommentInput) Validate() error {
	if c.ID == "" || c.UserID == "" {
		return errors.New("invalid or empty ids values")
	}
	return nil
}

func (c *UpdateCommentInput) ToComment(comment *Comment) {
	if c.Body != nil && comment.Body != *c.Body {
		comment.Body = *c.Body
	}
}
