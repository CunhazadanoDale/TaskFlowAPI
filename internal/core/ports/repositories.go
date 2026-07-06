package ports

import (
	"context"

	"github.com/CunhazadanoDale/TaskFlowAPI/internal/core/domain"
	"github.com/google/uuid"
)

type UserRepository interface {
	// Define methods for user repository
	CreateUser(ctx context.Context, user *domain.User) error
	GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

type ProjectRepository interface {
	// Define methods for project repository
	CreateProject(ctx context.Context, project *domain.Project) error
	GetProjectByID(ctx context.Context, id uuid.UUID) (*domain.Project, error)
	UpdateProject(ctx context.Context, project *domain.Project) error
	DeleteProject(ctx context.Context, id uuid.UUID) error
}

type TaskRepository interface {
	// Define methods for task repository
	CreateTask(ctx context.Context, task *domain.Task) error
	GetTaskByID(ctx context.Context, id uuid.UUID) (*domain.Task, error)
	GetTasksByAssigneeID(ctx context.Context, assigneeId uuid.UUID, filtro domain.PaginacaoFiltro) ([]*domain.Task, int, error)
	UpdateTask(ctx context.Context, task *domain.Task) error
	DeleteTask(ctx context.Context, id uuid.UUID) error
	GetTasksByProjectID(ctx context.Context, projectId uuid.UUID, filtro domain.PaginacaoFiltro) ([]*domain.Task, int, error)
}

type TaskHistoryRepository interface {
	// Define methods for task history repository
	CreateTaskHistory(ctx context.Context, history *domain.TaskHistory) error
	GetTaskHistoryByTaskID(ctx context.Context, taskId uuid.UUID, filtro domain.PaginacaoFiltro) ([]*domain.TaskHistory, int, error)
}