package domain

import (
	"errors"

	"github.com/google/uuid"
)

type TaskHistory struct {
	ID        string `json:"id" db:"id"`
	TaskId    uuid.UUID `json:"task_id" db:"task_id"`
	ChangedBy uuid.UUID `json:"changed_by" db:"changed_by"`
	Field    string `json:"field" db:"field"`
	OldValue string `json:"old_value" db:"old_value"`
	NewValue string `json:"new_value" db:"new_value"`
	ChangedAt int64  `json:"changed_at" db:"changed_at"`
}

func NewTaskHistory(taskId uuid.UUID, changedBy uuid.UUID, field string, oldValue string, newValue string) (*TaskHistory, error) {
	if (taskId == uuid.Nil) || (changedBy == uuid.Nil) || (field == "") {
		return nil, errors.New("taskId, changedBy and field are required")
	}

	return &TaskHistory{
		ChangedBy: changedBy,
		Field:     field,
		OldValue:  oldValue,
		NewValue:  newValue,
	}, nil
}
	