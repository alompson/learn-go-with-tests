package main

import "testing"

func TestHello(t *testing.T){
	t.Run("saying hello to people", func(t *testing.T){
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'hello, world' when an empty string is supplied", func(t *testing.T){
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say hello in Spanish", func(t *testing.T){
		got := Hello("Daniela", "Spanish")
		want := "Hola, Daniela"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say Hello World in Spanish", func(t *testing.T){
		got := Hello("", "Spanish")
		want := "Hola, Mundo"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say Hello in French", func(t *testing.T){
		got := Hello("Pierre", "French")
		want := "Bonjour, Pierre"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say Hello World in French", func(t *testing.T){
		got := Hello("", "French")
		want := "Bonjour, Monde"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want{
		t.Errorf("got %q want %q", got, want)
	}
}