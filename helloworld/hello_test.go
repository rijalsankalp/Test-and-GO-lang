package helloworld

import "testing"

func TestHello(t *testing.T) {
	t.Run("in English", func(t *testing.T) {
		got :=Hello("Kumar Krishna Tripathi", "English")
		want :="Hello, Kumar Krishna Tripathi"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got :=Hello("Kumar Krishna Tripathi", "Spanish")
		want :="Hola, Kumar Krishna Tripathi"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got :=Hello("Kumar Krishna Tripathi", "French")
		want :="Bonjour, Kumar Krishna Tripathi"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Nepali", func(t *testing.T) {
		got :=Hello("Kumar Krishna Tripathi", "Nepali")
		want :="Namaste, Kumar Krishna Tripathi"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World!' when an empty string is supplied", func(t *testing.T) {
		got :=Hello("", "")
		want :="Hello, World"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string){
	t.Helper()
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
