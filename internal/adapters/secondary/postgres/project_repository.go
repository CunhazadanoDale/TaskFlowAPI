package postgres

import (
	"context"
	"github.com/CunhazadanoDale/TaskFlowAPI/internal/core/domain"
	"github.com/CunhazadanoDale/TaskFlowAPI/internal/core/ports"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ ports.ProjectRepository = (*ProjectRepository)(nil)

type ProjectRepository struct {
	db *pgxpool.Pool
}

// CreateProject implements [ports.ProjectRepository].
func (p *ProjectRepository) CreateProject(ctx context.Context, project *domain.Project) error {
	query:= `INSERT INTO projects (id, name, description, owner_id, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := p.db.Exec(ctx, query, project.ID, project.Name, project.Description, project.OwnerId, project.CreatedAt, project.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

// DeleteProject implements [ports.ProjectRepository].
func (p *ProjectRepository) DeleteProject(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM projects WHERE id = $1`

	_, err := p.db.Exec(ctx, query, id)
	return err
}

// GetProjectByID implements [ports.ProjectRepository].
func (p *ProjectRepository) GetProjectByID(ctx context.Context, id uuid.UUID) (*domain.Project, error) {
	query := `SELECT id, name, description, owner_id, created_at, updated_at
	FROM projects WHERE id = $1`
	
	var project domain.Project
	err := p.db.QueryRow(ctx, query, id).Scan(&project.ID, &project.Name, &project.Description, &project.OwnerId, &project.CreatedAt, &project.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &project, nil
}

// UpdateProject implements [ports.ProjectRepository].
func (p *ProjectRepository) UpdateProject(ctx context.Context, project *domain.Project) error {
	query := `UPDATE projects SET name = $1, description = $2, updated_at = $3 WHERE id = $4`
	_, err := p.db.Exec(ctx, query, project.Name, project.Description, project.UpdatedAt, project.ID)
	return err
}

func NewProjectRepository(db *pgxpool.Pool) *ProjectRepository {
	return &ProjectRepository{db: db}
}
