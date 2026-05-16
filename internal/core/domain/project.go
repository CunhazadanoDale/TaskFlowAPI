package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Project struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	OwnerId     uuid.UUID      `json:"owner_id" db:"owner_id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

func NewProject(name string, description string, ownerId uuid.UUID) (*Project, error) {
	if name == "" {
		return nil, errors.New("project name is required")
	}
	if ownerId == uuid.Nil {
		return nil, errors.New("project owner ID is required")
	}

	return &Project{
		ID: uuid.New(),
		Name:        name,
		Description: description,
		OwnerId:     ownerId,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}
