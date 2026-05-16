package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Status      Status    `json:"status" db:"status"`
	Priority    Priority  `json:"priority" db:"priority"`
	DueDate     time.Time `json:"due_date" db:"due_date"`
	ProjectId   uuid.UUID `json:"project_id" db:"project_id"`
	CreatedBy   uuid.UUID `json:"created_by" db:"created_by"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type Status string

const (
	StatusPending    Status = "PENDING"
	StatusInProgress Status = "IN_PROGRESS"
	StatusCompleted  Status = "COMPLETED"
	StatusCancelled  Status = "CANCELLED"
)

type Priority string

const (
	PriorityLow    Priority = "LOW"
	PriorityMedium Priority = "MEDIUM"
	PriorityHigh   Priority = "HIGH"
)

func NewTask(title string, description string, status Status, priority Priority, dueDate time.Time, projectId uuid.UUID, createdBy uuid.UUID) (*Task, error) {
	if title == "" {
		return nil, errors.New("task title is required")
	}
	if projectId == uuid.Nil || createdBy == uuid.Nil {
		return nil, errors.New("project ID and created-by ID are required")
	}

	return &Task{
		Title: title,
		Description: description,
		Status: status,
		Priority: priority,
		DueDate: dueDate,
		ProjectId: projectId,
		CreatedBy: createdBy,
	}, nil
}