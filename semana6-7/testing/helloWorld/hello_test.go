package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)

	})

	t.Run("Say Hello in Spanish", func(t *testing.T) {
		got := Hello("Mario", "Spanish")
		want := "Hola, Mario"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say Hello in French", func(t *testing.T) {
		got := Hello("Mario", "French")
		want := "Bonjour, Mario"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Say Hello in Galician", func(t *testing.T) {
		got := Hello("Mario", "Galician")
		want := "Ola, Mario"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
