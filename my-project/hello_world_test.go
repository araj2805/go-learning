package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Ankit")
		want := "Hello Ankit "

		// if got != want {
		// 	t.Errorf("got %q want %q", got, want)
		// }

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying empty paramter", func(t *testing.T) {
		got := Hello("")
		want := "Hello World"

		// if got != want {
		// 	t.Errorf("got %q want %q", got, want)
		// }

		assertCorrectMessage(t, got, want)

	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got = %q, want = %q", got, want)
	}
}
