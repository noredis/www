package vo

import (
	"account-management-service/internal/failure"
	"golang.org/x/crypto/bcrypt"
	"unicode"
	"unicode/utf8"
)

type Password struct {
	value string
}

func NewPassword(password string, passwordConfirmation string) (*Password, error) {
	if password == "" {
		return nil, failure.EmptyPasswordError{}
	}

	if utf8.RuneCountInString(password) < 6 {
		return nil, failure.PasswordTooShortError{}
	}

	if utf8.RuneCountInString(password) > 40 {
		return nil, failure.PasswordTooLongError{}
	}

	if password != passwordConfirmation {
		return nil, failure.PasswordMismatchError{}
	}

	var (
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasUpper || !hasLower || !hasNumber || !hasSpecial {
		return nil, failure.InvalidPasswordError{}
	}

	password, err := hash(password)
	if err != nil {
		return nil, failure.UnableToHashPasswordError{}
	}

	return &Password{password}, nil
}

func hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (p Password) Value() string {
	return p.value
}

func (p Password) Compare(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(p.value), []byte(password))
	return err == nil
}
