package failure

type EmptyUsernameError struct{}

func (e EmptyUsernameError) Error() string {
	return "username should not be empty"
}

type UsernameTooShortError struct{}

func (e UsernameTooShortError) Error() string {
	return "username must be greater than 4"
}

type InvalidUsernameError struct{}

func (e InvalidUsernameError) Error() string {
	return "username can contain only lowercase Latin letters, numbers, underscores, and dots"
}

type UsernameTooLongError struct{}

func (e UsernameTooLongError) Error() string {
	return "username must be less than 20 characters"
}
