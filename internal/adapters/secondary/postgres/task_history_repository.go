package postgres

import (
	"context"
	"github.com/CunhazadanoDale/TaskFlowAPI/internal/core/domain"
	"github.com/CunhazadanoDale/TaskFlowAPI/internal/core/ports"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ ports.TaskHistoryRepository = (*TaskHistoryRepository)(nil)

type TaskHistoryRepository struct {
	db *pgxpool.Pool
}

// CreateTaskHistory implements [ports.TaskHistoryRepository].
func (t *TaskHistoryRepository) CreateTaskHistory(ctx context.Context, history *domain.TaskHistory) error {
	query := `INSERT INTO task_history (id, task_id, changed_by, field, old_value, new_value, changed_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := t.db.Exec(ctx, query, history.ID, history.TaskId, history.ChangedBy, history.Field, history.OldValue, history.NewValue, history.ChangedAt)
	return err
}

// GetTaskHistoryByTaskID implements [ports.TaskHistoryRepository].
func (t *TaskHistoryRepository) GetTaskHistoryByTaskID(ctx context.Context, taskId uuid.UUID, filtro domain.PaginacaoFiltro) ([]*domain.TaskHistory, int, error) {
	
	offset := (filtro.Page - 1) * filtro.PageSize

	query := `SELECT id, task_id, changed_by, field, old_value, new_value, changed_at
	FROM task_history WHERE task_id = $1 ORDER BY changed_at DESC LIMIT $2 OFFSET $3`
	
	var histories []*domain.TaskHistory
	rows, err := t.db.Query(ctx, query, taskId, filtro.PageSize, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	countQuery := `SELECT COUNT(*) FROM task_history WHERE task_id = $1`
	var count int
	err = t.db.QueryRow(ctx, countQuery, taskId).Scan(&count)
	if err != nil {
		return nil, 0, err
	}

	for rows.Next() {
		var history domain.TaskHistory
		err := rows.Scan(&history.ID, &history.TaskId, &history.ChangedBy, &history.Field, &history.OldValue, &history.NewValue, &history.ChangedAt)
		if err != nil {
			return nil, 0, err
		}
		histories = append(histories, &history)
	}

	return histories, count, nil
}

func NewTaskHistoryRepository(db *pgxpool.Pool) *TaskHistoryRepository {
	return &TaskHistoryRepository{db: db}
}
