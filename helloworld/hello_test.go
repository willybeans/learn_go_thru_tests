package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("run foreign: spanish", func(t *testing.T) {
		got := Hello("Maria", "Spanish")
		want := "Hola, Maria"
		assertCorrectMessage(t, got, want)
	})
	t.Run("run foreign: french", func(t *testing.T) {
		got := Hello("Maria", "French")
		want := "Bonjour, Francois"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	// t.Helper() tells our test to report the line numbers
	// from the function that called them
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}