package failure

type EmptyFullNameError struct{}

func (e EmptyFullNameError) Error() string {
	return "full name should not be empty"
}
