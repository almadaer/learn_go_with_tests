package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		// t.Helper() is needed to tell the test suite
		// that this method is a helper.
		// By doing this when it fails the line number
		// reported will be in our function call rather
		// than inside our test helper.
		t.Helper()
		if got != want {
			t.Errorf("got %q : want %q", got, want)
		}
	}
	t.Run("saying hello to people",
		func(t *testing.T) {
			got := hello("Joe", english)
			want := "Hello, Joe"
			assertCorrectMessage(t, got, want)
		})
	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := hello("", english)
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Spanish", func(t *testing.T) {
		got := hello("Elodie", spanish)
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("French", func(t *testing.T) {
		got := hello("Elodie", french)
		want := "Bonjour, Elodie"
		assertCorrectMessage(t, got, want)
	})
}
