package arraysandslices

import (
	"reflect"
	"testing"
)

func TestArrayAndSlicesSum(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
		t.Errorf("got '%v' want '%v'", got, want)
		}
	}
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := [5]int{1,2,3,4,6}
		got := ArraySum(numbers)
		want := 16

		if got != want {
		t.Errorf("got '%v' want '%v' given '%v'", got, want, numbers)
	}
	})
	t.Run("slices test", func(t *testing.T) {
		numbers := []int{1,2,3}
		got := SlicesSum(numbers)
		want := 6
		
		if got != want {
		t.Errorf("got '%v' want '%v' given '%v'", got, want, numbers)
	}
	})

	t.Run("SumAll test", func(t *testing.T) {
		got := SumAll([]int{1,2}, []int{0,9})
		want := []int{3,9}

		checkSums(t, got, want)
	})

	t.Run("Sum all tails", func(t *testing.T) {
		got := SumAllTails([]int{1,2}, []int{0,9})
		want := []int{2,9}

		checkSums(t, got, want)
	})

	t.Run("Sum all tails with empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0,9})
		want := []int{0,9}

		checkSums(t, got, want)
	})
}

