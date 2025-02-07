package failure

type EmptyEmailError struct{}

func (e EmptyEmailError) Error() string {
	return "email address should not be empty"
}

type InvalidEmailError struct{}

func (e InvalidEmailError) Error() string {
	return "invalid email address"
}
