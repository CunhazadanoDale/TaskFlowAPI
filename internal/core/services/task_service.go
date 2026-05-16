package services

import (
	"context"

	"github.com/CunhazadanoDale/TaskFlowAPI/internal/core/domain"
	"github.com/CunhazadanoDale/TaskFlowAPI/internal/core/ports"
	"github.com/google/uuid"
)

type TaskService struct {
	taskRepository        ports.TaskRepository
	projectRepository     ports.ProjectRepository
	taskHistoryRepository ports.TaskHistoryRepository
}

func NewTaskService(taskRepo ports.TaskRepository, projectRepo ports.ProjectRepository, taskHistoryRepo ports.TaskHistoryRepository) *TaskService {
	return &TaskService{
		taskRepository:        taskRepo,
		projectRepository:     projectRepo,
		taskHistoryRepository: taskHistoryRepo,
	}
}


func (s *TaskService) CreateTask(ctx context.Context, projectID uuid.UUID, title string, description string, status domain.Status) error {
	// Implement task creation logic here
	return nil
}

func (s *TaskService) GetTaskByID(ctx context.Context, taskID uuid.UUID) (*domain.Task, error) {
	// Implement task retrieval logic here
	return nil, nil
}

func (s *TaskService) GetTasksByProjectID(ctx context.Context, projectID uuid.UUID) ([]*domain.Task, error) {
	// Implement logic to retrieve tasks by project ID
	return nil, nil
}

func (s *TaskService) GetByAssigneeID(ctx context.Context, assigneeID uuid.UUID) ([]*domain.Task, error) {
	// Implement logic to retrieve tasks by assignee ID
	return nil, nil
}

func (s *TaskService) UpdateStatus(ctx context.Context, taskID uuid.UUID, newStatus domain.Status, changedBy uuid.UUID) error {
	// Implement task status update logic here, including validation and history tracking
	return nil
}

func (s *TaskService) DeleteTask(ctx context.Context, taskID uuid.UUID) error {
	// Implement task deletion logic here
	return nil
}

