package entity

import (
	"account-management-service/internal/domain/domainevent"
	vo2 "account-management-service/internal/domain/valueobject"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	fullName          vo2.FullName
	email             vo2.Email
	username          vo2.Username
	password          vo2.Password
	createdAt         time.Time
	passwordUpdatedAt time.Time
	emailConfirmedAt  *time.Time
	entity
}

func Register(
	id uuid.UUID,
	fullName vo2.FullName,
	email vo2.Email,
	username vo2.Username,
	password vo2.Password,
	now time.Time,
) *Account {
	account := &Account{
		entity:            entity{id: id},
		fullName:          fullName,
		email:             email,
		username:          username,
		password:          password,
		createdAt:         now,
		passwordUpdatedAt: now,
		emailConfirmedAt:  nil,
	}

	account.raiseDomainEvent(domainevent.NewAccountCreatedEvent(id, fullName, email, username))

	return account
}

func RestoreAccount(
	id uuid.UUID,
	fullName vo2.FullName,
	email vo2.Email,
	username vo2.Username,
	password vo2.Password,
	createdAt time.Time,
	passwordUpdatedAt time.Time,
	emailConfirmedAt *time.Time,
) *Account {
	return &Account{
		entity:            entity{id: id},
		fullName:          fullName,
		email:             email,
		username:          username,
		password:          password,
		createdAt:         createdAt,
		passwordUpdatedAt: passwordUpdatedAt,
		emailConfirmedAt:  emailConfirmedAt,
	}
}

func (a *Account) FullName() vo2.FullName {
	return a.fullName
}

func (a *Account) Email() vo2.Email {
	return a.email
}

func (a *Account) Username() vo2.Username {
	return a.username
}

func (a *Account) Password() vo2.Password {
	return a.password
}

func (a *Account) CreatedAt() time.Time {
	return a.createdAt
}

func (a *Account) PasswordUpdatedAt() time.Time {
	return a.passwordUpdatedAt
}

func (a *Account) EmailConfirmedAt() *time.Time {
	return a.emailConfirmedAt
}
