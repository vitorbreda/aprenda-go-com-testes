package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to someone", func(t *testing.T) {
		got := Hello("John", "")
		want := "Hello John"

		assertHello(t, got, want)
	})
	t.Run("saying hello world, when no one is specified", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World"

		assertHello(t, got, want)
	})
	t.Run("saying hello in Portuguese", func(t *testing.T) {
		got := Hello("João", "pt")
		want := "Olá João"

		assertHello(t, got, want)
	})

	t.Run("saying hello in Spanish", func(t *testing.T) {
		got := Hello("Juan", "es")
		want := "Hola Juan"

		assertHello(t, got, want)
	})
}

func assertHello(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
