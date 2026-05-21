package postgres

import (
	"context"
	"github.com/CunhazadanoDale/TaskFlowAPI/internal/core/domain"
	"github.com/CunhazadanoDale/TaskFlowAPI/internal/core/ports"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ ports.UserRepository = (*UserRepository)(nil)

type UserRepository struct {
	db *pgxpool.Pool
}

// CreateUser implements [ports.UserRepository].
func (u *UserRepository) CreateUser(ctx context.Context, user *domain.User) error {
	query := `INSERT INTO users (id, name, email, password_hash, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := u.db.Exec(ctx, query, user.ID, user.Name, user.Email, user.PasswordHash, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

// DeleteUser implements [ports.UserRepository].
func (u *UserRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM users WHERE id = $1`

	_, err := u.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

// GetUserByEmail implements [ports.UserRepository].
func (u *UserRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `SELECT id, name, email, password_hash, created_at, updated_at
	FROM users WHERE email = $1`

	var user domain.User
	row := u.db.QueryRow(ctx, query, email)

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByID implements [ports.UserRepository].
func (u *UserRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	query := `SELECT id, name, email, password_hash, created_at, updated_at
	FROM users WHERE id = $1`

	var user domain.User
	row := u.db.QueryRow(ctx, query, id)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser implements [ports.UserRepository].
func (u *UserRepository) UpdateUser(ctx context.Context, user *domain.User) error {
	query := `UPDATE users SET name = $1, email = $2, password_hash = $3, updated_at = $4
	WHERE id = $5`

	_, err := u.db.Exec(ctx, query, user.Name, user.Email, user.PasswordHash, user.UpdatedAt, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}
