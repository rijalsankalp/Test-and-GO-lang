package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("repeat 6 times", func(t *testing.T) {
	repeated := Repeat ("a", 6)
	expected := "aaaaaa"
	assertCorrectMessage(t, repeated, expected)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 7)
	}
}

func assertCorrectMessage(t testing.TB, repeated, expected string) {
	t.Helper()
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}