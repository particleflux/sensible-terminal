package sensible_terminal

type NotFoundError struct {
}

func (e *NotFoundError) Error() string {
	return "No known terminal emulator found"
}
