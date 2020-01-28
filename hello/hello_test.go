package main

import (
	"testing"
)

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Greating people", func(t *testing.T) {
		got := hello("Rian")
		want := "Hello Rian"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Hello Gorld when empty", func(t *testing.T) {
		got := hello("")
		want := "Hello Gorld"
		assertCorrectMessage(t, got, want)
	})

}
