package service

import (
	"github.com/google/uuid"
	"time"
)

type CreateAccountDTO struct {
	FullName             string
	Email                string
	Username             string
	Password             string
	PasswordConfirmation string
}

type AccountDTO struct {
	ID                uuid.UUID
	FullName          string
	Email             string
	Username          string
	CreatedAt         time.Time
	UpdatedAt         *time.Time
	PasswordUpdatedAt time.Time
	EmailConfirmedAt  *time.Time
}
