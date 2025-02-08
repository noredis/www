package data

import (
	"github.com/google/uuid"
	"time"
)

type AccountModel struct {
	ID                uuid.UUID
	FullName          string
	Email             string
	Username          string
	Password          string
	CreatedAt         time.Time
	PasswordUpdatedAt time.Time
	EmailConfirmedAt  *time.Time
}
