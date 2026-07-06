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
	panic("unimplemented")
}

// DeleteProject implements [ports.ProjectRepository].
func (p *ProjectRepository) DeleteProject(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}

// GetProjectByID implements [ports.ProjectRepository].
func (p *ProjectRepository) GetProjectByID(ctx context.Context, id uuid.UUID) (*domain.Project, error) {
	panic("unimplemented")
}

// UpdateProject implements [ports.ProjectRepository].
func (p *ProjectRepository) UpdateProject(ctx context.Context, project *domain.Project) error {
	panic("unimplemented")
}

func NewProjectRepository(db *pgxpool.Pool) *ProjectRepository {
	return &ProjectRepository{db: db}
}
