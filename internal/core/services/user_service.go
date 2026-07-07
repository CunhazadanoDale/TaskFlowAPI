package services

import (
	"context"
	"strings"
	"time"

	"github.com/CunhazadanoDale/TaskFlowAPI/internal/core/domain"
	"github.com/CunhazadanoDale/TaskFlowAPI/internal/core/ports"
	"github.com/google/uuid"
)

type UserService struct {
	// Define user-related methods here
	userRepository ports.UserRepository
}

func NewUserService(userRepo ports.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepo,
	}
}


func (s *UserService) RegisterUser(ctx context.Context, name string, email string, password string) error {
	nome := strings.TrimSpace(name)
	mail = strings.TrimSpace(email)
	senha = strings.TrimSpace(password)

	if nome == "" || mail == "" || senha == "" {
		return domain.ErrInvalidInput
	}

	now := time.Now()

	user := &domain.User{
		ID:           uuid.New(),
		Name:         nome,
		Email:        mail,
		PasswordHash: hashPassword(senha), // Implement password hashing
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	if userFromDB, _ := s.userRepository.GetUserByEmail(ctx, user); userFromDB != nil {
		return domain.ErrConflict
	}

	return s.userRepository.CreateUser(ctx, user)
	
}

func (s *UserService) GetUserByID(ctx context.Context, userID uuid.UUID) (*domain.User, error) {
	// Implement logic to retrieve a user by their ID
	return nil, nil
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	// Implement logic to retrieve a user by their email
	return nil, nil
}

func (s *UserService) AuthenticateUser(ctx context.Context, email string, password string) (*domain.User, error) {
	// Implement user authentication logic here, including password verification
	return nil, nil
} 

func (s *UserService) UpdateUser(ctx context.Context, userID uuid.UUID, name string, email string) error {
	// Implement logic to update user information
	return nil
}

func (s *UserService) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	// Implement logic to delete a user
	return nil
}