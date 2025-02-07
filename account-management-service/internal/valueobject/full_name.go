package vo

import "account-management-service/internal/failure"

type FullName struct {
	value string
}

func NewFullName(value string) (*FullName, error) {
	if value == "" {
		return nil, failure.EmptyFullNameError{}
	}

	return &FullName{value: value}, nil
}

func (f FullName) Value() string {
	return f.value
}
