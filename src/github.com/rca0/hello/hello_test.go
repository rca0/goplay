package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	assertCorretMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("foo", "")
		want := "Hello, foo"

		assertCorretMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorretMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("foo", "Spanish")
		want := "Hola, foo"

		assertCorretMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("foo", "French")
		want := "Bonjour, foo"

		assertCorretMessage(t, got, want)
	})
}

func ExampleHello() {
	fmt.Println("Hello")

	// Output: Hello
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Println("hello")
	}
}
