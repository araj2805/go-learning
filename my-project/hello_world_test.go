package main

import (
	"testing"
)

func TestHello(t *testing.T) {

	t.Run("Paramter Testing", func(t *testing.T) {
		got := Hello("Ankit", "")
		want := "Hello Ankit"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Empty Parameter", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Greeting in Spanish", func(t *testing.T) {
		got := Hello("Ankit", "Spanish")
		want := "Hola Ankit"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Greeting in French", func(t *testing.T) {
		got := Hello("Ankit", "French")
		want := "Bonjor Ankit"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got = %q, want = %q", got, want)
	}
}
