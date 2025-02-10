package domainevent

import (
	vo "account-management-service/internal/valueobject"
	"github.com/google/uuid"
)

type AccountCreatedEvent struct {
	accountID uuid.UUID
	fullName  vo.FullName
	email     vo.Email
	username  vo.Username
}

func NewAccountCreatedEvent(accountID uuid.UUID, fullName vo.FullName, email vo.Email, username vo.Username) *AccountCreatedEvent {
	return &AccountCreatedEvent{
		accountID: accountID,
		fullName:  fullName,
		email:     email,
		username:  username,
	}
}

func (e AccountCreatedEvent) GetEventName() string {
	return "EVENT.ACCOUNT_CREATED"
}

func (e AccountCreatedEvent) AccountID() uuid.UUID {
	return e.accountID
}

func (e AccountCreatedEvent) FullName() vo.FullName {
	return e.fullName
}

func (e AccountCreatedEvent) Email() vo.Email {
	return e.email
}

func (e AccountCreatedEvent) Username() vo.Username {
	return e.username
}
