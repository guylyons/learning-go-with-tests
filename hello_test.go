package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Guy")
	want := "Hello, Guy"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
