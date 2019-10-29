package sensible_terminal

import "testing"

func TestNotFoundError_Error(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		e := NotFoundError{}
		got := e.Error()
		want := "No known terminal emulator found"
		if got != want {
			t.Errorf("NotFoundError.Error() = %v, want %v", got, want)
		}
	})
}
