package entity

import (
	vo "account-management-service/internal/valueobject"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	id                uuid.UUID
	fullName          vo.FullName
	email             vo.Email
	username          vo.Username
	password          vo.Password
	createdAt         time.Time
	passwordUpdatedAt time.Time
	emailConfirmedAt  *time.Time
}

func Register(
	id uuid.UUID,
	fullName vo.FullName,
	email vo.Email,
	username vo.Username,
	password vo.Password,
	now time.Time,
) *Account {
	return &Account{
		id:                id,
		fullName:          fullName,
		email:             email,
		username:          username,
		password:          password,
		createdAt:         now,
		passwordUpdatedAt: now,
		emailConfirmedAt:  nil,
	}
}

func RestoreAccount(
	id uuid.UUID,
	fullName vo.FullName,
	email vo.Email,
	username vo.Username,
	password vo.Password,
	createdAt time.Time,
	passwordUpdatedAt time.Time,
	emailConfirmedAt *time.Time,
) *Account {
	return &Account{
		id:                id,
		fullName:          fullName,
		email:             email,
		username:          username,
		password:          password,
		createdAt:         createdAt,
		passwordUpdatedAt: passwordUpdatedAt,
		emailConfirmedAt:  emailConfirmedAt,
	}
}

func (a *Account) ID() uuid.UUID {
	return a.id
}

func (a *Account) FullName() vo.FullName {
	return a.fullName
}

func (a *Account) Email() vo.Email {
	return a.email
}

func (a *Account) Username() vo.Username {
	return a.username
}

func (a *Account) Password() vo.Password {
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
