package models

import (
	"errors"
	"strings"
)

type User struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Surname      string `json:"surname"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

type CreateUserInput struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LogUserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserAttributes struct {
	ID    string `json:"user_id"`
	Email string `json:"email"`
}

func (u *CreateUserInput) Validate() error {
	if u.Name == "" || u.Surname == "" {
		return errors.New("name fields can't be empty")
	}

	if !strings.Contains(u.Email, "@") {
		return errors.New("invalid email")
	}

	if len([]rune(u.Password)) < 8 {
		return errors.New("invalid password size")
	}

	return nil
}

func (u *LogUserInput) Validate() error {
	if !strings.Contains(u.Email, "@") {
		return errors.New("invalid email")
	}

	if len([]rune(u.Password)) < 8 {
		return errors.New("invalid password size")
	}

	return nil
}
