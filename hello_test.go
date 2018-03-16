package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessages := func(t *testing.T, got, want string) {
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Guy")
		want := "Hello, Guy"
		assertCorrectMessages(t, got, want)
	})

	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessages(t, got, want)
	})
}
