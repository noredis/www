package data

import (
	"account-management-service/internal/entity"
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
)

type DBClient interface {
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}

type AccountRepository struct {
	db DBClient
}

func NewAccountRepository(db DBClient) *AccountRepository {
	return &AccountRepository{db: db}
}

func (r AccountRepository) GetByID(ctx context.Context, id uuid.UUID) *entity.Account {
	return nil
}
func (r AccountRepository) GetByEmail(ctx context.Context, email string) *entity.Account {
	return nil
}

func (r AccountRepository) GetByUsername(ctx context.Context, username string) *entity.Account {
	return nil
}

func (r AccountRepository) Create(ctx context.Context, account entity.Account) error {
	const sql = `INSERT INTO users (id, full_name, email, username, password, created_at, password_updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := r.db.ExecContext(
		ctx,
		sql,
		account.ID(),
		account.FullName(),
		account.Email(),
		account.Username(),
		account.Password().Value(),
		account.CreatedAt(),
		account.PasswordUpdatedAt(),
	)
	if err != nil {
		const format = "failed insertion of Account to database: %v"
		return fmt.Errorf(format, err)
	}

	return nil
}
