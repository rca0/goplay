package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("foo")
	want := "Hello, foo"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
