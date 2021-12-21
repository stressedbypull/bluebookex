package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Davide")
	want := "Hello, Davide"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
