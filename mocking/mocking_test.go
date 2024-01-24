package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCountDown(t *testing.T) {
	t.Run("spy countdown", func(t *testing.T) {

		sleep := &SpySleeperOperations{}
		buffer := &bytes.Buffer{}

		CountDown(buffer, sleep)

		got := buffer.String()

		want := "3\n2\n1\nGo!"

		if got != want {
			t.Errorf("want %v, got %v", want, got)
		}

		if sleep.times != 3 {
			t.Errorf("the count of sleep invoked is %v, but the wanted values is 3.", sleep.times)
		}
	})

	t.Run("spy countdown operations", func(t *testing.T) {
		spysleep := &SpySleeperOperations{}
		CountDown(spysleep, spysleep)
		got := spysleep.Count

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %v, but got %v", want, got)
		}
	})
}