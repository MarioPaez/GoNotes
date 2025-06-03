package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Mario")
	want := "Hello Mario!"

	if got != want {
		t.Errorf("TestHello failed. Want: %s. Got: %s", want, got)
	}
}
