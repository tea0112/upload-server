package helper

import "testing"

func TestIsEmpty(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		got := IsEmpty("")
		want := true

		if got != want {
			t.Errorf("string is not empty")
		}
	})
	t.Run("not empty", func(t *testing.T) {
		got := IsEmpty("this is not empty")
		want := false

		if got != want {
			t.Errorf("string is empty")
		}
	})
}
