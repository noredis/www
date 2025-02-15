package service

import (
	"account-management-service/internal/domain/entity"
	"context"
)

type Publisher interface {
	Publish(message any) error
}

type AccountContextImp struct {
	repo      AccountRepository
	publisher Publisher
}

func NewAccountContextImp(repo AccountRepository, publisher Publisher) *AccountContextImp {
	return &AccountContextImp{repo: repo, publisher: publisher}
}

func (aci AccountContextImp) SaveChanges(ctx context.Context, account entity.Account) error {
	err := aci.repo.Create(ctx, account)
	if err != nil {
		return err
	}

	for _, e := range account.DomainEvents() {
		err = aci.publisher.Publish(e)
		if err != nil {
			return err
		}
	}

	return nil
}
