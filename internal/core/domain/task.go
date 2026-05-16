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
	DueDate     *time.Time `json:"due_date" db:"due_date"`
	ProjectId   uuid.UUID `json:"project_id" db:"project_id"`
	AssigneeId  *uuid.UUID `json:"assignee_id" db:"assignee_id"`
	CreatedBy   uuid.UUID `json:"created_by" db:"created_by"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type Status string

const (
	StatusPending    Status = "pending"
	StatusInProgress Status = "in_progress"
	StatusCompleted  Status = "completed"
	StatusCancelled  Status = "cancelled"
)

type Priority string

const (
	PriorityLow    Priority = "low"
	PriorityMedium Priority = "medium"
	PriorityHigh   Priority = "high"
)

func NewTask(title string, description string, status Status, priority Priority, dueDate time.Time, projectId uuid.UUID, assigneeId uuid.UUID, createdBy uuid.UUID) (*Task, error) {
	if title == "" {
		return nil, errors.New("task title is required")
	}
	if projectId == uuid.Nil || createdBy == uuid.Nil {
		return nil, errors.New("project ID and created-by ID are required")
	}

	return &Task{
		ID:		  uuid.New(),
		Title: title,
		Description: description,
		Status: status,
		Priority: priority,
		AssigneeId: &assigneeId,
		DueDate: &dueDate,
		ProjectId: projectId,
		CreatedBy: createdBy,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}


func (task *Task) CanBeEditedBy(userId uuid.UUID) bool {
	// Implement logic to check if the user has permission to edit the task
	if (task.AssigneeId != nil && *task.AssigneeId == userId) {
		return true

	}
	return userId == task.CreatedBy
}