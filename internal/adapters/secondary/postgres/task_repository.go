package postgres

import (
	"context"
	"github.com/CunhazadanoDale/TaskFlowAPI/internal/core/domain"
	"github.com/CunhazadanoDale/TaskFlowAPI/internal/core/ports"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ ports.TaskRepository = (*TaskRepository)(nil)

type TaskRepository struct {
	db *pgxpool.Pool
}

func NewTaskRepository(db *pgxpool.Pool) *TaskRepository {
	return &TaskRepository{db: db}
}

// CreateTask implements [ports.TaskRepository].
func (t *TaskRepository) CreateTask(ctx context.Context, task *domain.Task) error {
	query := `INSERT INTO tasks (id, title, description, status, priority, due_date, assignee_id, project_id, created_by, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err := t.db.Exec(ctx, query, task.ID, task.Title, task.Description, task.Status, task.Priority, task.DueDate, task.AssigneeId, task.ProjectId, task.CreatedBy, task.CreatedAt, task.UpdatedAt)
	return err
}

// DeleteTask implements [ports.TaskRepository].
func (t *TaskRepository) DeleteTask(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM tasks WHERE id = $1`
	_, err := t.db.Exec(ctx, query, id)
	return err
}

// GetTaskByID implements [ports.TaskRepository].
func (t *TaskRepository) GetTaskByID(ctx context.Context, id uuid.UUID) (*domain.Task, error) {
	query := `SELECT id, title, description, status, priority, due_date, assignee_id, project_id, created_by, created_at, updated_at FROM tasks WHERE id = $1`
	var task domain.Task
	err := t.db.QueryRow(ctx, query, id).Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.Priority, &task.DueDate, &task.AssigneeId, &task.ProjectId, &task.CreatedBy, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

// GetTasksByAssigneeID implements [ports.TaskRepository].
func (t *TaskRepository) GetTasksByAssigneeID(ctx context.Context, assigneeId uuid.UUID, filtro domain.PaginacaoFiltro) ([]*domain.Task, int, error) {
	offset := (filtro.Page - 1) * filtro.PageSize

	query := `SELECT id, title, description, status, priority, due_date, assignee_id, project_id, created_by, created_at, updated_at FROM tasks WHERE assignee_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`
	var tasks []*domain.Task
	rows, err := t.db.Query(ctx, query, assigneeId, filtro.PageSize, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	countQuery := `SELECT COUNT(*) FROM tasks WHERE assignee_id = $1`
	var count int
	err = t.db.QueryRow(ctx, countQuery, assigneeId).Scan(&count)
	if err != nil {
		return nil, 0, err
	}

	for rows.Next() {
		var task domain.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.Priority, &task.DueDate, &task.AssigneeId, &task.ProjectId, &task.CreatedBy, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return nil, 0, err
		}
		tasks = append(tasks, &task)
	}

	return tasks, count, nil
}

// GetTasksByProjectID implements [ports.TaskRepository].
func (t *TaskRepository) GetTasksByProjectID(ctx context.Context, projectId uuid.UUID, filtro domain.PaginacaoFiltro) ([]*domain.Task, int, error) {
	offset := (filtro.Page - 1) * filtro.PageSize

	query := `SELECT id, title, description, status, priority, due_date, assignee_id, project_id, created_by, created_at, updated_at FROM tasks WHERE project_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`
	var tasks []*domain.Task
	rows, err := t.db.Query(ctx, query, projectId, filtro.PageSize, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	queryCount := `SELECT COUNT(*) FROM tasks WHERE project_id = $1`
	var count int
	err = t.db.QueryRow(ctx, queryCount, projectId).Scan(&count)
	if err != nil {
		return nil, 0, err
	}

	for rows.Next() {
		var task domain.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.Priority, &task.DueDate, &task.AssigneeId, &task.ProjectId, &task.CreatedBy, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return nil, 0, err
		}
		tasks = append(tasks, &task)
	}

	return tasks, count, nil
}

// UpdateTask implements [ports.TaskRepository].
func (t *TaskRepository) UpdateTask(ctx context.Context, task *domain.Task) error {
	panic("unimplemented")
}
