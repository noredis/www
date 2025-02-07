package vo

import (
	"account-management-service/internal/failure"
	"net/mail"
)

type Email struct {
	value string
}

func NewEmail(value string) (*Email, error) {
	if value == "" {
		return nil, failure.EmptyEmailError{}
	}

	addr, err := mail.ParseAddress(value)
	if err != nil {
		return nil, failure.InvalidEmailError{}
	}

	return &Email{value: addr.Address}, nil
}

func (e Email) Value() string {
	return e.value
}
