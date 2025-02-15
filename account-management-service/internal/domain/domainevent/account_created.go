package domainevent

import (
	vo2 "account-management-service/internal/domain/valueobject"
	"github.com/google/uuid"
)

type AccountCreatedEvent struct {
	accountID uuid.UUID
	fullName  vo2.FullName
	email     vo2.Email
	username  vo2.Username
}

func NewAccountCreatedEvent(accountID uuid.UUID, fullName vo2.FullName, email vo2.Email, username vo2.Username) *AccountCreatedEvent {
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

func (e AccountCreatedEvent) FullName() vo2.FullName {
	return e.fullName
}

func (e AccountCreatedEvent) Email() vo2.Email {
	return e.email
}

func (e AccountCreatedEvent) Username() vo2.Username {
	return e.username
}
