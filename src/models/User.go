package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// omitempty - If field is empty then the property will be omitted
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// Prepare will call validate and formatt methods to validate and format the received user data
func (user *User) Prepare(stage string) error {
	if erro := user.validate(stage); erro != nil {
		return erro
	}

	if erro := user.format(stage); erro != nil {
		return erro
	}

	return nil
}

func (user *User) validate(stage string) error {
	if user.Name == "" {
		return errors.New("Name is required and cannot be empty")
	}

	if user.Nick == "" {
		return errors.New("Nick is required and cannot be empty")
	}

	if user.Email == "" {
		return errors.New("Email is required and cannot be empty")
	}

	if erro := checkmail.ValidateFormat(user.Email); erro != nil {
		return errors.New("Email is invalid")
	}

	if stage == "signup" && user.Password == "" {
		return errors.New("Password is required and cannot be empty")
	}

	return nil
}

func (user *User) format(stage string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if stage == "signup" {
		hashPassword, erro := security.Hash(user.Password)
		if erro != nil {
			return erro
		}

		user.Password = string(hashPassword)
	}

	return nil
}
