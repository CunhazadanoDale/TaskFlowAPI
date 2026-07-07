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
	mail := strings.TrimSpace(email)
	senha := strings.TrimSpace(password)

	if nome == "" || mail == "" || senha == "" {
		return domain.ErrInvalidInput
	}

	if _, err := s.userRepository.GetUserByEmail(ctx, mail); err == nil {
		return domain.ErrConflict
	}

	user, err := domain.NewUser(nome, mail, HashPassword(senha))
	if err != nil {
		return err
	}

	return s.userRepository.CreateUser(ctx, user)
	
}

func (s *UserService) GetUserByID(ctx context.Context, userID uuid.UUID) (*domain.User, error) {
	user, err := s.userRepository.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	user, err := s.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) AuthenticateUser(ctx context.Context, email string, password string) (*domain.User, error) {
	mail := strings.TrimSpace(email)
	senha := strings.TrimSpace(password)

	if mail == "" || senha == "" {
		return nil, domain.ErrInvalidInput
	}

	user, err := s.userRepository.GetUserByEmail(ctx, mail)
	if err != nil {
		return nil, err
	}

	if !CheckPasswordHash(senha, user.PasswordHash) { // Implement password hash checking
		return nil, domain.ErrUnauthorized
	}

	return user, nil
}

func (s *UserService) UpdateUser(ctx context.Context, userID uuid.UUID, name string, email string) error {
	nome := strings.TrimSpace(name)
	mail := strings.TrimSpace(email)

	if nome == "" || mail == "" {
		return domain.ErrInvalidInput
	}

	user, err := s.userRepository.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}

	user.Name = nome
	user.Email = mail
	user.UpdatedAt = time.Now()

	return s.userRepository.UpdateUser(ctx, user)
}

func (s *UserService) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	return s.userRepository.DeleteUser(ctx, userID)
}