package vo

import (
	"account-management-service/internal/core/failure"
)

type FullName struct {
	value string
}

func NewFullName(fullName string) (*FullName, error) {
	if fullName == "" {
		return nil, failure.EmptyFullNameError{}
	}

	return &FullName{value: fullName}, nil
}

func (f FullName) Value() string {
	return f.value
}
