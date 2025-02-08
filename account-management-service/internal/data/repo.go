package data

import (
	"account-management-service/internal/entity"
	vo "account-management-service/internal/valueobject"
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

func (r AccountRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.Account, error) {
	const sql = `SELECT id, full_name, email, username, password, created_at, password_updated_at, email_confirmed_at FROM users WHERE id = $1`

	row := r.db.QueryRowContext(ctx, sql, id)

	var account AccountModel

	err := row.Scan(
		&account.ID,
		&account.FullName,
		&account.Email,
		&account.Username,
		&account.Password,
		&account.CreatedAt,
		&account.PasswordUpdatedAt,
		&account.EmailConfirmedAt,
	)
	if err != nil {
		const format = "failed selection of Account from database: %v"
		return nil, fmt.Errorf(format, err)
	}

	fullName, err := vo.NewFullName(account.FullName)
	if err != nil {
		return nil, err
	}

	email, err := vo.NewEmail(account.Email)
	if err != nil {
		return nil, err
	}

	username, err := vo.NewUsername(account.Username)
	if err != nil {
		return nil, err
	}

	password, err := vo.RestorePassword(account.Password)
	if err != nil {
		return nil, err
	}

	a := entity.RestoreAccount(
		account.ID,
		*fullName,
		*email,
		*username,
		*password,
		account.CreatedAt,
		account.PasswordUpdatedAt,
		account.EmailConfirmedAt,
	)

	return a, nil
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
		account.FullName().Value(),
		account.Email().Value(),
		account.Username().Value(),
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
