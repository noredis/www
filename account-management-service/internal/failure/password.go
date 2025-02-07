package failure

type PasswordTooShortError struct{}

func (e PasswordTooShortError) Error() string {
	return "password should be at least 6 characters"
}

type PasswordTooLongError struct{}

func (e PasswordTooLongError) Error() string {
	return "password should be at most 40 characters"
}

type PasswordMismatchError struct{}

func (e PasswordMismatchError) Error() string {
	return "password mismatch"
}

type InvalidPasswordError struct{}

func (e InvalidPasswordError) Error() string {
	return "password must contain an uppercase letter, a lowercase letter, a number, and a special character"
}

type EmptyPasswordError struct{}

func (e EmptyPasswordError) Error() string {
	return "password should not be an empty string"
}

type UnableToHashPasswordError struct{}

func (e UnableToHashPasswordError) Error() string {
	return "unable to hash password"
}
