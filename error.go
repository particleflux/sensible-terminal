package sensibleterminal

// NotFoundError is returned when no known terminal emulator can be found
type NotFoundError struct {
}

func (e *NotFoundError) Error() string {
	return "No known terminal emulator found"
}
