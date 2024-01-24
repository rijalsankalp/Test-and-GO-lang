package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("testing string putput", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Greet(&buffer, "Hello, World!")

		got := buffer.String()

		want := "Hello, World!"

		if got != want {
			t.Errorf(" want %s, got %s", want, got)
		}
	})
}