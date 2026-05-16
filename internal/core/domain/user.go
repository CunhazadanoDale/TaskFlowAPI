package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID    uuid.UUID `json:"id" db:"id"`
	Name  string    `json:"name" db:"name"`
	Email string    `json:"email" db:"email"`
	Password string    `json:"password" db:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}


func NewUser(name string, email string, password string) (*User, error) {
	if name == "" {
		return nil, errors.New("user name is required")
	}
	if email == "" {
		return nil, errors.New("user email is required")
	}
	if password == "" {
		return nil, errors.New("user password is required")
	}

	return &User{
		Name: name,
		Email: email,
		Password: password,
	}, nil
}