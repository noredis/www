package service

import (
	"account-management-service/internal/entity"
	"account-management-service/internal/failure"
	vo "account-management-service/internal/valueobject"
	"context"
	"github.com/google/uuid"
	"time"
)

type AccountRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Account, error)
	GetByEmail(ctx context.Context, email string) (*entity.Account, error)
	GetByUsername(ctx context.Context, username string) (*entity.Account, error)
	Create(ctx context.Context, account entity.Account) error
}

type Publisher interface {
	Publish(message any) error
}

type AccountService struct {
	accountRepository AccountRepository
	publisher         Publisher
}

func NewAccountService(accountRepository AccountRepository, publisher Publisher) *AccountService {
	return &AccountService{accountRepository: accountRepository, publisher: publisher}
}

func (acs *AccountService) CreateAccount(ctx context.Context, dto CreateAccountDTO) (*AccountDTO, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	fullName, err := vo.NewFullName(dto.FullName)
	if err != nil {
		return nil, err
	}

	email, err := vo.NewEmail(dto.Email)
	if err != nil {
		return nil, err
	}

	existing, err := acs.accountRepository.GetByEmail(ctx, email.Value())
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, failure.EmailIsBusyError{}
	}

	username, err := vo.NewUsername(dto.Username)
	if err != nil {
		return nil, err
	}

	existing, err = acs.accountRepository.GetByUsername(ctx, username.Value())
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, failure.UsernameIsBusyError{}
	}

	password, err := vo.NewPassword(dto.Password, dto.PasswordConfirmation)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	account := entity.Register(id, *fullName, *email, *username, *password, now)

	err = acs.accountRepository.Create(ctx, *account)
	if err != nil {
		return nil, err
	}

	for _, e := range account.DomainEvents() {
		err = acs.publisher.Publish(e)
		if err != nil {
			return nil, err
		}
	}

	return &AccountDTO{
		ID:                account.ID(),
		FullName:          account.FullName().Value(),
		Email:             account.Email().Value(),
		Username:          account.Username().Value(),
		CreatedAt:         account.CreatedAt(),
		UpdatedAt:         nil,
		PasswordUpdatedAt: account.PasswordUpdatedAt(),
		EmailConfirmedAt:  nil,
	}, nil
}
