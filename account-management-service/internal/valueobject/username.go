package vo

import (
	"account-management-service/internal/failure"
	"regexp"
	"unicode/utf8"
)

type Username struct {
	value string
}

func NewUsername(username string) (*Username, error) {
	if username == "" {
		return nil, failure.EmptyUsernameError{}
	}

	if utf8.RuneCountInString(username) < 5 {
		return nil, failure.UsernameTooShortError{}
	}

	if utf8.RuneCountInString(username) > 20 {
		return nil, failure.UsernameTooLongError{}
	}

	pattern := regexp.MustCompile(`^[a-z0-9_.]+$`)
	if !pattern.MatchString(username) {
		return nil, failure.InvalidUsernameError{}
	}

	return &Username{value: username}, nil
}

func (u Username) Value() string {
	return u.value
}
